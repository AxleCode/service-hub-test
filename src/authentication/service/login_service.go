package service

import (
	"context"
	"database/sql"
	"encoding/base64"


	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/common/jwt"
	"gitlab.com/wit-id/service-hub-test/common/utility"
	"gitlab.com/wit-id/service-hub-test/src/repository/payload"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
)

func (s *AuthenticationService) Login(ctx context.Context, request payload.LoginPayload, jwtRequest jwt.RequestJWTToken)(
	u sqlc.GetAuthenticationByUsernameRow, err error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil{
		log.FromCtx(ctx).Error(err, "failed begin tx")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	query := sqlc.New(s.mainDB).WithTx(tx)

	defer func ()  {
		if err != nil{
			if rollBackErr := tx.Rollback(); rollBackErr != nil{
				log.FromCtx(ctx).Error(err, "error rollback", rollBackErr)
				err = errors.WithStack(httpservice.ErrUnknownSource)

				return
			}
		}
	}()

	u, err = query.GetAuthenticationByUsername(ctx, request.Username)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get user")
		err = errors.WithStack(httpservice.ErrUserNotFound)
		return
	}

	// Assume u.Salt is a base64-encoded string
	storedSalt, err := base64.StdEncoding.DecodeString(u.Salt.String)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to decode salt")
		err = errors.WithStack(httpservice.ErrInternalServerError)
		return
	}

	// Hash the input password using the decoded salt
	hashedInputPassword := utility.HashPasswordLaravel(request.Password, storedSalt)
	if hashedInputPassword != u.Password {
		log.FromCtx(ctx).Error(err, "password do not match")
		err = errors.WithStack(httpservice.ErrPasswordNotMatch)
		return
	}

	if u.Status != "active" {
		log.FromCtx(ctx).Error(err, "user is not active")
		err = errors.WithStack(httpservice.ErrUserNotFound)
		return
	}

	// Update Last login user backoffice
	if err = query.RecordAuthenticationLastLogin(ctx, u.Guid); err != nil {
		log.FromCtx(ctx).Error(err, "failed record last login")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	// Update token auth record
	if err = query.RecordAuthTokenUserLogin(ctx, sqlc.RecordAuthTokenUserLoginParams{
		UserLogin: sql.NullString{
			String: u.Guid,
			Valid:  true,
		},
		Name:       jwtRequest.AppName,
		DeviceID:   jwtRequest.DeviceID,
		DeviceType: jwtRequest.DeviceType,
	}); err != nil {
		log.FromCtx(ctx).Error(err, "failed update token auth login user")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}