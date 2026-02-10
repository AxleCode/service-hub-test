package echokit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
	// "gitlab.com/wit-id/service-hub-test/toolkit/runtimekit"
)

type RuntimeConfig struct {
	Port                    int
	Name                    string
	BuildInfo               string
	ShutdownWaitDuration    time.Duration
	ShutdownTimeoutDuration time.Duration
	RequestTimeoutConfig    *TimeoutConfig
	HealthCheckPath         string
	InfoCheckPath           string
	HealthCheckFunc         func(ctx context.Context) error
}

type TimeoutConfig struct {
	Timeout time.Duration
	Skipper middleware.Skipper
}

const (
	defaultPort       = 8080
	defaultHealthPath = "/actuator/health"
	defaultInfoPath   = "/actuator/info"
	defaultReqTimeout = 7 * time.Second
)

func (cfg *RuntimeConfig) validate() {
	if cfg.Port == 0 {
		cfg.Port = defaultPort
	}
	if cfg.HealthCheckPath == "" {
		cfg.HealthCheckPath = defaultHealthPath
	}
	if cfg.RequestTimeoutConfig == nil {
		cfg.RequestTimeoutConfig = &TimeoutConfig{Timeout: defaultReqTimeout}
	}
	if cfg.RequestTimeoutConfig.Timeout == 0 {
		cfg.RequestTimeoutConfig.Timeout = defaultReqTimeout
	}
	if cfg.RequestTimeoutConfig.Skipper == nil {
		cfg.RequestTimeoutConfig.Skipper = middleware.DefaultSkipper
	}
	if cfg.InfoCheckPath == "" {
		cfg.InfoCheckPath = defaultInfoPath
	}
}

type healthStatus struct {
	Status string `json:"status"`
}

func RunServerWithContext(appCtx context.Context, e *echo.Echo, cfg *RuntimeConfig) {
	logger := log.FromCtx(appCtx)
	cfg.validate()

	hs := &healthStatus{Status: "UP"}

	e.GET(cfg.HealthCheckPath, func(c echo.Context) error {
		if cfg.HealthCheckFunc != nil {
			if err := cfg.HealthCheckFunc(c.Request().Context()); err != nil {
				return c.JSON(http.StatusOK, &healthStatus{Status: "OUT_OF_SERVICE"})
			}
		}
		return c.JSON(http.StatusOK, hs)
	})

	e.GET(cfg.InfoCheckPath, func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"version": cfg.BuildInfo})
	})

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout:  cfg.RequestTimeoutConfig.Timeout,
		Skipper:  cfg.RequestTimeoutConfig.Skipper,
	}))

	go func() {
		<-appCtx.Done()
		hs.Status = "OUT_OF_SERVICE"
		logger.Info("shutting down HTTP server")
		<-time.After(cfg.ShutdownWaitDuration)
		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeoutDuration)
		defer cancel()
		_ = e.Shutdown(shutdownCtx)
	}()

	logger.Info("serving REST HTTP server", "port", cfg.Port)
	if err := e.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error(err, "starting http server")
	}
}