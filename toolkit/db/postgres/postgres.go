package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/toolkit/db"
)

func NewPostgresDatabase(opt *db.Option) (*sql.DB, error) {
	port := opt.Port
	if port == 0 {
		port = 5432
	}
	connURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(opt.Username, opt.Password),
		Host:   fmt.Sprintf("%s:%d", opt.Host, port),
		Path:   opt.DatabaseName,
	}
	q := connURL.Query()
	q.Add("sslmode", "disable")
	connURL.RawQuery = q.Encode()

	sqlDB, err := sql.Open("postgres", connURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "postgres: failed to open connection")
	}

	sqlDB.SetMaxIdleConns(opt.ConnectionOption.MaxIdle)
	sqlDB.SetConnMaxLifetime(opt.ConnectionOption.MaxLifetime)
	sqlDB.SetMaxOpenConns(opt.ConnectionOption.MaxOpen)

	ctx, cancel := context.WithTimeout(context.Background(), opt.ConnectionOption.ConnectTimeout)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "postgres: ping failed")
	}

	log.Println("successfully connected to postgres", opt.Host+":"+strconv.Itoa(port))

	go keepAlive(sqlDB, opt.DatabaseName, opt.KeepAliveCheckInterval)

	return sqlDB, nil
}

func keepAlive(db *sql.DB, dbName string, interval time.Duration) {
	for {
		if _, err := db.Exec("SELECT 1"); err != nil {
			log.Printf("db keepalive error db_name=%s err=%v", dbName, err)
			return
		}
		time.Sleep(interval)
	}
}