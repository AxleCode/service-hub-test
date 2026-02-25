package application

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/src/check_weather/service"
	"gitlab.com/wit-id/service-hub-test/toolkit/config"
	"gitlab.com/wit-id/service-hub-test/src/repository/payload"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/src/middleware"

)

func AddRouteCheckWeather(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewCheckWeatherService(s.GetDB(), cfg)
	mddw := middleware.NewEnsureToken(s.GetDB(), cfg)

	checkWeatherApp := e.Group("/check-weather")

	checkWeatherApp.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Check Weather ok")
	})

	var client http.Client
	if cfg.GetBool("local-dev") {
		proxyUrl, err := url.Parse("http://127.0.0.1:8080")
		if err != nil {
		}

		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}

		client.Transport = transport
	}

	//untuk validasi token
	checkWeatherApp.Use(mddw.ValidateToken)
	//untuk validasi login
	checkWeatherApp.Use(mddw.ValidateUserLogin)

	checkWeatherApp.GET("/info", GetCheckWeather(svc, client, cfg))
}

func GetCheckWeather(svc *service.CheckWeatherService, client http.Client, cfg config.KVStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request payload.CheckWeatherPayload
		if err := c.Bind(&request); err != nil {
			log.FromCtx(c.Request().Context()).Error(err, "failed parse request body")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(); err != nil {
			return err
		}

		body, err := svc.GetCheckWeather(c.Request().Context(), client, cfg, request)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(c, body, nil)
	}
}
