package echohttp

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/wit-id/service-hub-test/common/constants"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/toolkit/config"
	"gitlab.com/wit-id/service-hub-test/toolkit/echokit"

	checkWeather "gitlab.com/wit-id/service-hub-test/src/check_weather/application"
)

func RunEchoHTTPService(ctx context.Context, s *httpservice.Service, cfg config.KVStore) {
	e := echo.New()
	e.HTTPErrorHandler = handleEchoError(cfg)
	e.Use(handleLanguage)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, constants.DefaultAllowHeaderAuthorization},
	}))

	runtimeCfg := echokit.NewRuntimeConfig(cfg, "restapi")
	runtimeCfg.HealthCheckFunc = s.GetServiceHealth

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	// add route here
	checkWeather.AddRouteCheckWeather(s, cfg, e)

	// end of route
	
	echokit.RunServerWithContext(ctx, e, runtimeCfg)
}