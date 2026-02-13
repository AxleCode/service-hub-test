package service

import (
	"context"

	"gitlab.com/wit-id/service-hub-test/common/utility"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
)

func (s *InventoryService) DeleteInventory(ctx context.Context, guid string) (err error) {
	responseData, err := utility.Transaction(ctx, s.mainDB, func(query *sqlc.Queries) (response interface{}, err error){
		response, err = query.UpdateStatusInventory(ctx, guid)
		return
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed delete Inventory")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return	
	}
	_ = responseData
	return
}