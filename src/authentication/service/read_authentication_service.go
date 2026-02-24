package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
)

func (s *AuthenticationService) ReadAuthenticationByID(ctx context.Context, requestGuid string) (authToken sqlc.GetAuthenticationByIDRow, err error){
	q := sqlc.New(s.mainDB)

	authToken, err = q.GetAuthenticationByID(ctx, requestGuid)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get authentication by phone number")
		err = errors.WithStack(httpservice.ErrInvalidToken)

		return
	}

	return
}