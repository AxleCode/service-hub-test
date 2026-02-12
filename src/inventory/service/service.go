package service

import (
	"database/sql"

	"gitlab.com/wit-id/service-hub-test/toolkit/config"
)

type InventoryService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewInventoryService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *InventoryService {
	return &InventoryService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}