package postgres

import (
	"database/sql"
	// "fmt"
	"strconv"

	"gitlab.com/wit-id/service-hub-test/toolkit/config"
	"gitlab.com/wit-id/service-hub-test/toolkit/db"
)

func NewFromConfig(cfg config.KVStore, path string) (*sql.DB, error) {
	connOpt := db.DefaultConnectionOption()

	if maxIdle := cfg.GetInt(path + ".conn.max-idle"); maxIdle > 0 {
		connOpt.MaxIdle = maxIdle
	}
	if maxOpen := cfg.GetInt(path + ".conn.max-open"); maxOpen > 0 {
		connOpt.MaxOpen = maxOpen
	}
	if maxLifetime := cfg.GetDuration(path + ".conn.max-lifetime"); maxLifetime > 0 {
		connOpt.MaxLifetime = maxLifetime
	}
	if keepAlive := cfg.GetDuration(path + ".conn.keep-alive-interval"); keepAlive > 0 {
		connOpt.KeepAliveCheckInterval = keepAlive
	}

	port, _ := strconv.Atoi(cfg.GetString(path + ".port"))
	if port == 0 {
		port = 5432
	}

	opt, err := db.NewDatabaseOption(
		cfg.GetString(path+".host"),
		port,
		cfg.GetString(path+".username"),
		cfg.GetString(path+".password"),
		cfg.GetString(path+".schema"),
		connOpt,
	)
	if err != nil {
		return nil, err
	}

	return NewPostgresDatabase(opt)
}