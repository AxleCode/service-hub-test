package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
)

func (s *InventoryService) GetInventory(ctx context.Context, key string) (
	post sqlc.GetOneInventoryRow, err error) {
	query := sqlc.New(s.mainDB)

	result, err := query.GetOneInventory(ctx, key)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get barang")
		err = errors.WithStack(httpservice.ErrDataNotFound)
		return
	}

	return result, nil
}

func (s *InventoryService) ListInventory(ctx context.Context, request sqlc.ListInventoryParams) (
	listInventory []sqlc.ListInventoryRow, totalData int64, err error) {
	query := sqlc.New(s.mainDB)

	//Get Total Data
	totalData, err = s.getTotalListInventory(ctx, query, request)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get total list inventory")
		err = errors.WithStack(httpservice.ErrDataNotFound)
		return
	}

	listInventory, err = query.ListInventory(ctx, request)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get list inventory")
		err = errors.WithStack(httpservice.ErrDataNotFound)
		return
	}

	return
}

func (s *InventoryService) getTotalListInventory(ctx context.Context, query *sqlc.Queries, request sqlc.ListInventoryParams) (
	totalData int64, err error) {
		requestParam := sqlc.CountListInventoryParams{
		BarangID:    request.BarangID,
		Jumlah:    request.Jumlah,
		Status:    request.Status,
	}
	totalData, err = query.CountListInventory(ctx, requestParam)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get total list inventory")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}
	return
}