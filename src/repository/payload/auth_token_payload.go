package payload

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
)

type AuthTokenPayload struct {
	AppName	string `json:"app_name" validate:"required"`
	AppKey	string `json:"app_key" validate:"required"`
	DeviceID	string `json:"device_id" validate:"required"`
	DeviceType	string `json:"device_type" validate:"required"`
	IPAddress	string `json:"ip_address" validate:"required"`
}

type readAuthTokenPayload struct {
	Name string `json:"name"`
	DeviceID string `json:"device_id"`
	DeviceType string `json:"device_type"`
	Token string `json:"token"`
	TokenExpired time.Time `json:"token_expired"`
	RefreshToken string `json:"refresh_token"`
	RefreshTokenExpired time.Time `json:"refresh_token_expired"`
	IsLogin bool `json:"is_login"`
	UserLogin string `json:"user_login"`
}

// Validate validates the AuthTokenPayload fields
func (p *AuthTokenPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(p); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	
	return
}

//ToPayloadAuthToken
func ToPayloadAuthToken(data sqlc.AuthenticationSchemaAuthToken) (response readAuthTokenPayload){
	response = readAuthTokenPayload{
		Name : data.Name,
		DeviceID: data.DeviceID,
		DeviceType: data.DeviceType,
		Token: data.Token,
		TokenExpired: data.TokenExpired,
		RefreshToken: data.RefreshToken,
		RefreshTokenExpired: data.RefreshTokenExpired,
		IsLogin: data.IsLogin,
	}

	if data.UserLogin.Valid {
		response.UserLogin = data.UserLogin.String
	}

	return
}