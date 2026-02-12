package service

import (
	"database/sql"

	"gitlab.com/wit-id/service-hub-test/toolkit/config"
)

type BarangService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewBarangService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *BarangService {
	return &BarangService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
