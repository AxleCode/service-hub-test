package constants

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gojektech/heimdall/httpclient"
	"github.com/pkg/errors"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/service-hub-test/toolkit/config"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
)

const (
	AccessView     = "view"
	AccessCreate   = "create"
	AccessUpdate   = "update"
	AccessPut      = "put"
	AccessApproval = "approval"
	AccessGet      = "get"
	AccessList     = "list"
	AccessDelete   = "delete"

	ConfigRoutes                 = "common.config-routes-key"
	ConfigPrefixRoutesBackoffice = "common.prefix-config-route-backoffice"

	MddwTokenKey       = "token-data"
	MddwUserData       = "user-data"
	MddwUserHandheld   = "user-handheld"
	MddwUserBackoffice = "user-backoffice"
	MddwUserAccess     = "user-access"
	MddwKeyRole        = "role-data"

	ConfigAppointment     = "config_appointment"
	ConfigFeeRegisPatient = "config_fee_registration_patient"

	SuperadminConstantGuid    = "c4ae1c7a-d1aa-4667-bd54-6f4d396a84d6"
	SeksiPastoralFollowUpGuid = "5a8d006c-9508-40d0-9a20-1092522da0f8"

	PengurusFamilyAltarIDGembala    = "33213d5c-8828-44ee-893b-0dfc1f49f14f"
	PengurusFamilyAltarIDWakil      = "59bfcd9d-d242-4a9f-8d4d-d5984a43bf75"
	PengurusFamilyAltarIDSekretaris = "3f5e0a4f-0274-46d2-a5d3-5c31b8ab3fe8"

	TimeExpire    = 24
	timeoutNumber = 20000
	retryNumber   = 2
)

type HttpRequestPayload struct {
	Method         string
	PathURL        string
	Body           *[]byte
	HeaderOptional []HeaderOptional
}

// HeaderOptional ...
type HeaderOptional struct {
	Key   string
	Value string
}

func WordpressHttpRequest(q *sqlc.Queries, cfg config.KVStore, request HttpRequestPayload) (body []byte, err error) {
	timeout := timeoutNumber * time.Millisecond
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		httpclient.WithRetryCount(retryNumber),
	)

	var newReq *http.Request
	var reqBody []byte
	if request.Body != nil {
		reqBody = *request.Body
		newReq, _ = http.NewRequest(request.Method, cfg.GetString("wordpress-api.host")+request.PathURL, bytes.NewBuffer(reqBody))
	} else {
		newReq, _ = http.NewRequest(request.Method, cfg.GetString("wordpress-api.host")+request.PathURL, nil)
	}

	newReq.Header.Add("Content-Type", "application/json")

	// if request have header optional
	if len(request.HeaderOptional) > 0 {
		for _, v := range request.HeaderOptional {
			newReq.Header[v.Key] = []string{v.Value}
			//newReq.Header.Add(v.Key, v.Value)
		}
	}

	// Generate and log the curl command
	curlCommand := generateCurlCommand(newReq, reqBody)
	log.FromCtx(context.Background()).Info("Executing curl command: " + curlCommand)

	response, err := client.Do(newReq)
	if err != nil {
		log.FromCtx(context.Background()).Error(err, "failed send request to powerpro api")
		err = errors.Wrap(err, "")

		return
	}

	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		errNew := errors.New(string(body))
		log.FromCtx(context.Background()).Error(errNew, "failed send request to powerpro api")
		err = errors.Wrap(errNew, "")

		return
	}

	return
}

func generateCurlCommand(req *http.Request, reqBody []byte) string {
	curlCommand := fmt.Sprintf("curl -X %s '%s'", req.Method, req.URL.String())

	// Add headers
	for key, values := range req.Header {
		for _, value := range values {
			curlCommand += fmt.Sprintf(" -H '%s: %s'", key, value)
		}
	}

	// Add body if present
	if len(reqBody) > 0 {
		curlCommand += fmt.Sprintf(" -d '%s'", string(reqBody))
	}

	return curlCommand
}
