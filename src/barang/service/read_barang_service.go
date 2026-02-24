package service

import (
	"context"
	

	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	// "gitlab.com/wit-id/service-hub-test/src/repository/payload"

)

func (s *BarangService) GetBarang(ctx context.Context, key string) (
	post sqlc.GetOneBarangRow, err error) {
	query := sqlc.New(s.mainDB)

	result, err := query.GetOneBarang(ctx, key)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get barang")
		err = errors.WithStack(httpservice.ErrDataNotFound)
		return
	}

	// response := payload.BarangResponse{
	// 	Guid:       result.Guid,
	// 	KodeBarang: result.KodeBarang,
	// 	NamaBarang: result.NamaBarang,
	// 	Deskripsi:  result.Deskripsi.String,
	// 	Kategori:   result.Kategori,
	// 	Harga:      result.Harga,
	// }

	return result, nil
}

func (s *BarangService) ListBarang(ctx context.Context, request sqlc.ListBarangParams) (
	listBarang []sqlc.ListBarangRow, totalData int64, err error) {
	query := sqlc.New(s.mainDB)

	//Get Total Data
	totalData, err = s.getTotalListBarang(ctx, query, request)
	if err != nil {
		return
	}

	listBarang, err = query.ListBarang(ctx, request)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get list barang")
		err = errors.WithStack(httpservice.ErrDataNotFound)
		return
	}

	return
}

func (s *BarangService) getTotalListBarang(ctx context.Context, query *sqlc.Queries, request sqlc.ListBarangParams) (
	totalData int64, err error) {
		requestParam := sqlc.CountListBarangParams{
			Kategori: request.Kategori,
			SetKategori: request.SetKategori,
			SetKodeBarang: request.SetKodeBarang,
			KodeBarang: request.KodeBarang,
			SetNamaBarang: request.SetNamaBarang,
			NamaBarang: request.NamaBarang,
			SetGuid: request.SetGuid,
			Guid: request.Guid,
		}

	totalData, err = query.CountListBarang(ctx, requestParam)
	if err != nil {
		log.FromCtx(ctx).Error(err, "Failed get total data Barang")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}
