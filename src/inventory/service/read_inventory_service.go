package service

import (
	"context"
	// "database/sql"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
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
		BarangID: request.BarangID,
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

func (s *InventoryService) CountStockProduct(ctx context.Context, barangID string) (
	sqlc.CountInventoryStokByBarangIDRow, error) {
    query := sqlc.New(s.mainDB)

    data, err := query.CountInventoryStokByBarangID(ctx, barangID)
    if err != nil {
        log.FromCtx(ctx).Error(err, "failed count stock product")
        return sqlc.CountInventoryStokByBarangIDRow{}, 
            errors.WithStack(httpservice.ErrUnknownSource)
    }

    return data, nil
}

func (s *InventoryService) CountAllStockProduct(ctx context.Context, request sqlc.ListCountAllInventoryEachProductParams) (
	listCountAllInventory []sqlc.ListCountAllInventoryEachProductRow, totalData int64, err error) {
	query := sqlc.New(s.mainDB)

	//Get Total Data
	totalData, err = s.getTotalCountAllInventory(ctx, query, request)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get total count all inventory")
		err = errors.WithStack(httpservice.ErrDataNotFound)
		return
	}

	listCountAllInventory, err = query.ListCountAllInventoryEachProduct(ctx, request)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed count all stock product")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}

	log.FromCtx(ctx).Info("filter debug",
		"set_nama_barang", request.SetNamaBarang,
		"nama_barang", request.NamaBarang,
		"set_kategori", request.SetKategori,
		"kategori", request.Kategori,
	)
	// log.FromCtx(ctx).Info("list inventory", "data", listCountAllInventory)

	return
}

func (s *InventoryService) getTotalCountAllInventory(
		ctx context.Context,
		query *sqlc.Queries,
		request sqlc.ListCountAllInventoryEachProductParams,
	) (int64, error) {
	countReq := sqlc.CountListCountAllInventoryEachProductParams{
		SetNamaBarang: request.SetNamaBarang,
		NamaBarang:    request.NamaBarang,
		SetKategori:   request.SetKategori,
		Kategori:      request.Kategori,
	}
	totalData, err := query.CountListCountAllInventoryEachProduct(
		ctx,
		countReq,
	)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get total count all inventory")
		return 0, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return totalData, nil
}
