package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/common/utility"
	"gitlab.com/wit-id/service-hub-test/src/repository/payload"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
)

func (s *InventoryService) CreateInventory(ctx context.Context, request sqlc.InsertInventoryParams, p payload.InsertInventoryPayload) (
	 inventoryGUID string, err error) {
		q := sqlc.New(s.mainDB)

		inventoryData, err := q.InsertInventory(ctx, request)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed to insert inventory", err)
			err = errors.Wrap(httpservice.ErrBadRequestWithMessage,
				utility.ParseSqlError(err))
			return
		}

		inventoryGUID = inventoryData.Guid
		return 
}