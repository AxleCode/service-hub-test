package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	// "gitlab.com/wit-id/service-hub-test/src/repository/payload"
	"gitlab.com/wit-id/service-hub-test/common/utility"
)

func (s *BarangService) UpdateBarang(ctx context.Context, data sqlc.UpdateBarangParams) (
	post sqlc.GetOneBarangRow, err error) {
		
		responseData, err := utility.Transaction(ctx, s.mainDB, func(query *sqlc.Queries) (response interface{}, err error){
			response, err = query.UpdateBarang(ctx, data)
			return
		})
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed update Barang")
			err = errors.WithStack(httpservice.ErrUnknownSource)

			return	
		}
		//Get response data post update
		post, err = s.GetBarang(ctx, responseData.(sqlc.Barang).Guid)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed get Barang after update")
			err = errors.WithStack(httpservice.ErrUnknownSource)
			return
		}
		return
	}