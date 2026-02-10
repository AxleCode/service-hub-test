package main

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gitlab.com/wit-id/service-hub-test/common/echohttp"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/toolkit/db/postgres"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
	"gitlab.com/wit-id/service-hub-test/toolkit/runtimekit"
	"google.golang.org/grpc"
	"gitlab.com/wit-id/service-hub-test/toolkit/grpckit"
)

func main() {
	var err error
	setDefaultTimezone()

	appContext, cancel := runtimekit.NewRuntimeContext()
	defer cancel()
	defer func() {
		if err != nil {
			log.FromCtx(appContext).Error(err, "application error")
		}
	}()

	cfg, err := loadConfig("config.yaml")
	if err != nil {
		return
	}

	mainDB, err := postgres.NewFromConfig(cfg, "db")
	if err != nil {
		return
	}
	defer mainDB.Close()

	logger, err := log.NewFromConfig(cfg, "log")
	if err != nil {
		return
	}
	logger.Set()

	// Initialize HTTP service
	svc := httpservice.NewService(mainDB, cfg)

	// Run gRPC server
	grpcSrv := grpc.NewServer()
	grpcCfg := grpckit.NewRuntimeConfig(cfg, "grpc")
	grpcCfg.HealthCheckFunc = svc.GetServiceHealth
	go grpckit.RunWithContext(appContext, grpcSrv, grpcCfg)
	
	// Run HTTP server
	echohttp.RunEchoHTTPService(appContext, svc, cfg)
}

func setDefaultTimezone() {
	loc, _ := time.LoadLocation("UTC")
	time.Local = loc
}

func loadConfig(filePath string) (*viper.Viper, error) {
	cfg := viper.New()
	cfg.SetConfigFile(filePath)
	if err := cfg.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "reading config")
	}
	return cfg, nil
}