package service

import (
	"database/sql"

	"gitlab.com/wit-id/service-hub-test/toolkit/config"
)

type AuthenticationService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewAuthenticationService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *AuthenticationService {
	return &AuthenticationService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
