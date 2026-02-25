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
	barang "gitlab.com/wit-id/service-hub-test/src/barang/application"
	inventory "gitlab.com/wit-id/service-hub-test/src/inventory/application"
	authToken "gitlab.com/wit-id/service-hub-test/src/auth_token/application"
	authentication "gitlab.com/wit-id/service-hub-test/src/authentication/application"
)

func RunEchoHTTPService(ctx context.Context, s *httpservice.Service, cfg config.KVStore) {
	e := echo.New()
	e.HTTPErrorHandler = handleEchoError(cfg)
	e.Use(handleLanguage)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, constants.DefaultAllowHeaderToken, constants.DefaultAllowHeaderRefreshToken, constants.DefaultAllowHeaderAuthorization},
	}))

	runtimeCfg := echokit.NewRuntimeConfig(cfg, "restapi")
	runtimeCfg.HealthCheckFunc = s.GetServiceHealth

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	// add route here
	checkWeather.AddRouteCheckWeather(s, cfg, e)
	barang.AddRouteBarang(s, cfg, e)
	inventory.AddRouteInventory(s, cfg, e)
	authToken.AddRouteAuthToken(s, cfg, e)
	authentication.AddRouteAuthentication(s, cfg, e)

	// end of route
	
	echokit.RunServerWithContext(ctx, e, runtimeCfg)
}