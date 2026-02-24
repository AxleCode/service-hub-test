package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/common/utility"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
)

func (s *AuthenticationService) AuthenticationRegister(ctx context.Context, request sqlc.InsertAuthenticationParams) (response sqlc.AuthenticationSchemaAuthentication, err error){
	var errParsing bool
	tx, err := s.mainDB.BeginTx(ctx, nil)
	if err != nil {
		log.FromCtx(ctx).Error(err, "error init Transaction", err)
		err = errors.Wrap(httpservice.ErrBadRequestWithMessage,
			utility.ParseSqlError(err))
		return
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback", rollBackErr)
				err = errors.WithStack(httpservice.ErrUnknownSource)

				return
			}
			if errParsing {
				err = nil
			}
			return
		}
		tx.Commit()
	}()

	response, err = q.InsertAuthentication(ctx, request)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed register User, please check you already put the correct payload")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}

	return
}

func (s *AuthenticationService) UserRegisterFromWeb(ctx context.Context, request sqlc.RegisterUserParams) (response sqlc.AuthenticationSchemaUser, err error) {
	var errParsing bool
	tx, err := s.mainDB.BeginTx(ctx, nil)
	if err != nil {
		log.FromCtx(ctx).Error(err, "error Init Transaction", err)
		err = errors.Wrap(httpservice.ErrBadRequestWithMessage,
			utility.ParseSqlError(err))
		return
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback", rollBackErr)
				err = errors.WithStack(httpservice.ErrUnknownSource)

				return
			}
			if errParsing {
				err = nil
			}
			return
		}
		tx.Commit()
	}()

	response, err = q.RegisterUser(ctx, request)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed register User, please check you already put the correct payload")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}

	return
}