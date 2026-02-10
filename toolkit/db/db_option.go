package db

import (
	"time"
	"github.com/pkg/errors"
)

const (
	defaultMaxOpen           = 100
	defaultMaxLifetime       = 10 * time.Minute
	defaultMaxIdle           = 5
	defaultConnectTimeout    = 10 * time.Second
	defaultKeepAliveInterval = 30 * time.Second
)

type Option struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
	*ConnectionOption
}

type ConnectionOption struct {
	MaxIdle                int
	MaxLifetime            time.Duration
	MaxOpen                int
	ConnectTimeout         time.Duration
	KeepAliveCheckInterval time.Duration
}

func DefaultConnectionOption() *ConnectionOption {
	return &ConnectionOption{
		MaxIdle:                5,
		MaxOpen:                100,
		MaxLifetime:            10 * time.Minute,
		ConnectTimeout:         10 * time.Second,
		KeepAliveCheckInterval: 30 * time.Second,
	}
}

var errInvalidDBSource = errors.New("invalid datasource host | port")

func NewDatabaseOption(host string, port int, username, password, dbName string, conn *ConnectionOption) (*Option, error) {
	if host == "" || port == 0 {
		return nil, errors.Wrapf(errInvalidDBSource, "db: host=%s port=%d", host, port)
	}
	if conn == nil || conn.MaxOpen == 0 {
		conn = DefaultConnectionOption()
	}
	return &Option{
		Host:             host,
		Port:             port,
		Username:         username,
		Password:         password,
		DatabaseName:     dbName,
		ConnectionOption: conn,
	}, nil
}