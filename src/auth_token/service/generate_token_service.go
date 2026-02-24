package service

import (
	"database/sql"
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/common/jwt"
	"gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/service-hub-test/src/repository/payload"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
	
)

func (s *AuthTokenService) AuthToken(ctx context.Context, request payload.AuthTokenPayload) (authToken sqlc.AuthenticationSchemaAuthToken, err error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin transaction")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback transaction", rollBackErr)
				err = errors.WithStack(httpservice.ErrUnknownSource)
				return
			}
		}
	}()

	// validate app key
	if err = s.validateAppKey(ctx, q, payload.ValidateAppKeyPayload{
		AppName: request.AppName,
		AppKey:  request.AppKey,
	}); err != nil {
		log.FromCtx(ctx).Error(err, "failed to validate app key")
		err = errors.WithStack(httpservice.ErrUnauthorizedTokenData)
		return
	}

	// generate token
	jwtResponse, err := jwt.CreateJWTToken(ctx, s.cfg, jwt.RequestJWTToken{
		AppName: request.AppName,
		DeviceID: request.DeviceID,
		DeviceType: request.DeviceType,
		IPAddress: request.IPAddress,
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to generate token")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}
		
	// log.FromCtx(ctx).Info("succeed to generate token", jwtResponse.Token)
	// log.FromCtx(ctx).Info("refresh token", jwtResponse.RefreshToken)

	authToken, err = s.recordToken(ctx, q, jwtResponse, false)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to record token")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}
	
	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "failed to commit transaction")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}
	
	return
}

func (s *AuthTokenService) validateAppKey(ctx context.Context, q *sqlc.Queries, request payload.ValidateAppKeyPayload) (err error) {
	appKeyData, err := q.GetAppKeyByName(ctx, request.AppName)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get app key")
		err = errors.WithStack(httpservice.ErrUnauthorizedTokenData)
		return
	}

	if appKeyData.Key != request.AppKey {
		log.FromCtx(ctx).Error(errors.New("invalid app key"), "invalid app key")
		err = errors.WithStack(httpservice.ErrUnauthorizedTokenData)
		return
	}

	return
}

func (s *AuthTokenService) recordToken(ctx context.Context, q *sqlc.Queries, token jwt.ResponseJwtToken, isRefreshToken bool) (authToken sqlc.AuthenticationSchemaAuthToken, err error) {
	if !isRefreshToken {
		authToken, err = q.InsertAuthToken(ctx, sqlc.InsertAuthTokenParams{
			Name: token.AppName,
			DeviceID: token.DeviceID,
			DeviceType: token.DeviceType,
			Token: token.Token,
			TokenExpired: token.TokenExpired,
			IpAddress: sql.NullString{
				String: token.IPAddress,
				Valid: true,
			},
			RefreshToken: token.RefreshToken,
			RefreshTokenExpired: token.RefreshTokenExpired,
		})
		
	} else {
		// Get record
		authData, errGetRecord := s.ReadAuthToken(ctx, sqlc.GetAuthTokenParams{
			Name: token.AppName,
			DeviceID: token.DeviceID,
			DeviceType: token.DeviceType,
		})
		if errGetRecord != nil {
			log.FromCtx(ctx).Error(errGetRecord, "failed to get auth token record")
			err = errors.WithStack(httpservice.ErrUnknownSource)
			return
		}

		authToken, err = q.InsertAuthToken(ctx, sqlc.InsertAuthTokenParams{
			Name: authData.Name,
			DeviceID: authData.DeviceID,
			DeviceType: authData.DeviceType,
			Token: token.Token,
			TokenExpired: token.TokenExpired,
			IpAddress: sql.NullString{
				String: token.IPAddress,
				Valid: true,
			},
			RefreshToken: token.RefreshToken,
			RefreshTokenExpired: token.RefreshTokenExpired,
			IsLogin: authData.IsLogin,
			UserLogin: authData.UserLogin,
		})

	}
	
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to insert auth token")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}
	
	return
}

func (s *AuthTokenService) RefreshToken(ctx context.Context, request jwt.RequestJWTToken) (authToken sqlc.AuthenticationSchemaAuthToken, err error){
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin tx")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func(){
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback=%s", rollBackErr)
				err = errors.WithStack(httpservice.ErrUnknownSource)

				return
			}
		}
	}()

	jwtResponse, err := jwt.CreateJWTToken(ctx, s.cfg, request)
	if err != nil {
		return
	}

	authToken, err = s.recordToken(ctx, q, jwtResponse, true)
	if err != nil{
		return
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}