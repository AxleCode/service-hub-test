package service

import (
	"database/sql"

	"gitlab.com/wit-id/service-hub-test/toolkit/config"
)

type AuthTokenService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewAuthTokenService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *AuthTokenService {
	return &AuthTokenService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}