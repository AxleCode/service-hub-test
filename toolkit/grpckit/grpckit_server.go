package grpckit

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/iancoleman/strcase"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const (
	defaultPort                = 8288
	defaultReqTimeout          = 7 * time.Second
	defaultShutdownWaitTimeout = 7 * time.Second
)

type RuntimeConfig struct {
	Port                 int
	Name                 string
	RequestTimeout       time.Duration
	ShutdownWaitDuration time.Duration
	EnableReflection     bool
	HealthCheckFunc      HealthCheckFunc
}

func (cfg *RuntimeConfig) validate() {
	if cfg.Port == 0 {
		cfg.Port = defaultPort
	}
	if cfg.RequestTimeout == 0 {
		cfg.RequestTimeout = defaultReqTimeout
	}
	if cfg.ShutdownWaitDuration == 0 {
		cfg.ShutdownWaitDuration = defaultShutdownWaitTimeout
	}
}

func RunWithContext(appCtx context.Context, s *grpc.Server, cfg *RuntimeConfig) {
	cfg.Name = strcase.ToSnake(cfg.Name)
	cfg.validate()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.FromCtx(appCtx).Error(err, "grpc net.Listen", "port", cfg.Port)
		return
	}

	hs := NewHealthcheckServer(cfg.HealthCheckFunc)
	grpc_health_v1.RegisterHealthServer(s, hs)

	if cfg.EnableReflection {
		reflection.Register(s)
	}

	log.FromCtx(appCtx).Info("serving gRPC server", "port", cfg.Port)

	go func() {
		<-appCtx.Done()
		hs.Serving = false
		log.FromCtx(appCtx).Info("shutting down gRPC server", "wait_ms", cfg.ShutdownWaitDuration.Milliseconds())
		<-time.After(cfg.ShutdownWaitDuration)
		s.GracefulStop()
	}()

	if err := s.Serve(lis); err != nil {
		log.FromCtx(appCtx).Error(err, "grpc Serve")
	}
}