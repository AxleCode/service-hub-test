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

func (s *InventoryService) UpdateInventory(ctx context.Context, data sqlc.UpdateInventoryParams) (
	post sqlc.GetOneInventoryRow, err error) {
		
		responseData, err := utility.Transaction(ctx, s.mainDB, func(query *sqlc.Queries) (response interface{}, err error){
			response, err = query.UpdateInventory(ctx, data)
			return
		})
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed update Inventory")
			err = errors.WithStack(httpservice.ErrUnknownSource)

			return	
		}
		//Get response data post update
		post, err = s.GetInventory(ctx, responseData.(sqlc.Inventory).Guid)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed get Inventory after update")
			err = errors.WithStack(httpservice.ErrUnknownSource)
			return
		}
		return
	}