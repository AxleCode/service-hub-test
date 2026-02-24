package payload

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/utility"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
)

type RegisterPayload struct{
	Username        string `json:"username" valid:"required"`
	FirstName       string `json:"first_name" valid:"required"`
	LastName        string `json:"last_name" valid:"required"`
	PhoneNumber		string `json:"phone_number"`
	Email           string `json:"email" valid:"required"`
	Address			string `json:"address"`
	Password        string `json:"password" valid:"required"`
	ConfirmPassword string `json:"confirm_password" valid:"required"`
}

type readAuthentication struct {
	Guid 			string `json:"guid"`
	UserID 			string `json:"user_id"`
	UserName 		string `json:"username"`
	UserEmail 		string `json:"email"`
	UserPhoneNumber string `json:"phone_number"`
	UserFCMToken	string `json:"fcm_token"`
	Status			string `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedBy       string     `json:"created_by"`
	UpdatedAt       time.Time `json:"updated_at"`
	UpdatedBy       string    `json:"updated_by"`
}

type LoginPayload struct{
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

func (payload *RegisterPayload) Validate(ctx context.Context) (err error) {
	if err = utility.ValidateStruct(ctx, payload); err != nil {
		return
	}

	if payload.Password != payload.ConfirmPassword {
		return errors.New("password and confirm password not match")
	}

	return
}

func (payload *LoginPayload) Validate(ctx context.Context) (err error){
	if err = utility.ValidateStruct(ctx, payload); err != nil{
		return
	}
	return
}

func (payload *RegisterPayload) ToEntity() (data sqlc.RegisterUserParams){
	data = sqlc.RegisterUserParams{
		Guid:	utility.GenerateGoogleUUID(),
		Name: payload.FirstName + " " + payload.LastName,
		Email: payload.Email,
		Address: sql.NullString{
			String: payload.Address,
			Valid:  payload.Address != "",
		},
		PhoneNumber: payload.PhoneNumber,
	}

	return
}

func ToPayloadUserProfile(data sqlc.GetAuthenticationByIDRow) (response readAuthentication){
	response = readAuthentication{
		Guid: data.Guid,
		Status: data.Status,

		CreatedBy: data.CreatedBy,
	}
	if data.UserGuid.Valid {
		response.UserID = data.UserGuid.String
	}

	if data.Username.Valid {
		response.UserName = data.Username.String
	}

	if data.Email.Valid {
		response.UserEmail = data.Email.String
	}

	if data.FcmToken.Valid {
		response.UserFCMToken = data.FcmToken.String
	}

	if data.PhoneNumber.Valid {
		response.UserPhoneNumber = data.PhoneNumber.String
	}

	if data.CreatedAt.Valid {
		response.CreatedAt = data.CreatedAt.Time
	}

	if data.UpdatedAt.Valid {
		response.UpdatedAt = data.UpdatedAt.Time
	}

	if data.UpdatedBy.Valid {
		response.UpdatedBy = data.UpdatedBy.String
	}

	return
}

func ToPayloadAuthentication(data sqlc.GetAuthenticationByUsernameRow) (response readAuthentication) {
	response = readAuthentication{
		Guid:      data.Guid,
		Status:    data.Status,
		CreatedBy: data.CreatedBy,
	}

	if data.UserGuid.Valid {
		response.UserID = data.UserGuid.String
	}

	if data.Username.Valid {
		response.UserName = data.Username.String
	}

	if data.Email.Valid {
		response.UserEmail = data.Email.String
	}

	if data.FcmToken.Valid {
		response.UserFCMToken = data.FcmToken.String
	}

	if data.PhoneNumber.Valid {
		response.UserPhoneNumber = data.PhoneNumber.String
	}

	if data.CreatedAt.Valid {
		response.CreatedAt = data.CreatedAt.Time
	}

	if data.UpdatedAt.Valid {
		response.UpdatedAt = data.UpdatedAt.Time
	}

	if data.UpdatedBy.Valid {
		response.UpdatedBy = data.UpdatedBy.String
	}

	return
}