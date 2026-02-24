package smtp

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/toolkit/config"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
)

// Attachment represents an email attachment
type Attachment struct {
	Filename    string
	ContentType string
	Data        []byte
}

func SendMail(ctx context.Context, to []string, subject, message string, cfg config.KVStore) error {
	return SendMailWithAttachments(ctx, to, subject, message, nil, cfg)
}

// SendMailWithAttachments sends an email with optional attachments
func SendMailWithAttachments(ctx context.Context, to []string, subject, message string, attachments []Attachment, cfg config.KVStore) error {
	boundary := "boundary123456789"

	// Build email body with attachments
	body := "From: " + cfg.GetString("smtp.sender") + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: multipart/mixed; boundary=" + boundary + "\n\n"

	// Add HTML message part
	body += "--" + boundary + "\n"
	body += "Content-Type: text/html; charset=UTF-8\n\n"
	body += message + "\n\n"

	// Add attachments if any
	for _, attachment := range attachments {
		body += "--" + boundary + "\n"
		body += "Content-Type: " + attachment.ContentType + "\n"
		body += "Content-Disposition: attachment; filename=\"" + attachment.Filename + "\"\n"
		body += "Content-Transfer-Encoding: base64\n\n"

		// Encode attachment data as base64
		encodedData := base64.StdEncoding.EncodeToString(attachment.Data)
		// Split into lines of 76 characters (RFC 2045)
		for i := 0; i < len(encodedData); i += 76 {
			end := i + 76
			if end > len(encodedData) {
				end = len(encodedData)
			}
			body += encodedData[i:end] + "\n"
		}
		body += "\n"
	}

	// Close boundary
	body += "--" + boundary + "--\n"

	auth := smtp.PlainAuth("", cfg.GetString("smtp.auth_email"), cfg.GetString("smtp.auth_password"), cfg.GetString("smtp.host"))
	port, err := strconv.Atoi(cfg.GetString("smtp.port"))
	if err != nil {
		log.FromCtx(ctx).Error(err, "invalid smtp.port")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}
	smtpAddr := fmt.Sprintf("%s:%d", cfg.GetString("smtp.host"), port)

	err = smtp.SendMail(smtpAddr, auth, cfg.GetString("smtp.auth_email"), to, []byte(body))
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed send email")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return err
	}

	return nil
}

func SendTLSMail(ctx context.Context, to []string, subject, message string, cfg config.KVStore) error {
	body := "From: " + cfg.GetString("smtp.sender") + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n" +
		"Content-Type: text/html; charset=UTF-8\n\n" +
		message

	// Connect via TLS to port 465
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         cfg.GetString("smtp.host"),
	}

	port, err := strconv.Atoi(cfg.GetString("smtp.port"))
	if err != nil {
		log.FromCtx(ctx).Error(err, "invalid smtp.port")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", cfg.GetString("smtp.host"), port), tlsConfig)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to dial SMTP over TLS")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	client, err := smtp.NewClient(conn, cfg.GetString("smtp.host"))
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to create SMTP client")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	// Authenticate
	auth := smtp.PlainAuth("", cfg.GetString("smtp.auth_email"), cfg.GetString("smtp.auth_password"), cfg.GetString("smtp.host"))
	if err := client.Auth(auth); err != nil {
		log.FromCtx(ctx).Error(err, "SMTP auth failed")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err := client.Mail(cfg.GetString("smtp.auth_email")); err != nil {
		log.FromCtx(ctx).Error(err, "SMTP MAIL FROM failed")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	for _, addr := range to {
		if err := client.Rcpt(addr); err != nil {
			log.FromCtx(ctx).Error(err, "SMTP RCPT TO failed")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}
	}

	writer, err := client.Data()
	if err != nil {
		log.FromCtx(ctx).Error(err, "SMTP DATA failed")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	_, err = writer.Write([]byte(body))
	if err != nil {
		log.FromCtx(ctx).Error(err, "SMTP body write failed")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	err = writer.Close()
	if err != nil {
		log.FromCtx(ctx).Error(err, "SMTP data close failed")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	client.Quit()

	return nil
}
