package service

import (
	"context"

	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
)

func (s *AuthTokenService) ReadAuthToken(ctx context.Context, request sqlc.GetAuthTokenParams) (authToken sqlc.AuthenticationSchemaAuthToken, err error) {
	q := sqlc.New(s.mainDB)

	authToken, err = q.GetAuthToken(ctx, request)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get auth token")
		err = errors.WithStack(httpservice.ErrDataNotFound)
		return
	}

	return
}