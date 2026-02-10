package service

import (
	"database/sql"

	"gitlab.com/wit-id/service-hub-test/toolkit/config"
)

type CheckWeatherService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewCheckWeatherService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *CheckWeatherService {
	return &CheckWeatherService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
