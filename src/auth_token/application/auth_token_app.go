package application

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/common/jwt"
	"gitlab.com/wit-id/service-hub-test/src/auth_token/service"
	"gitlab.com/wit-id/service-hub-test/src/middleware"
	"gitlab.com/wit-id/service-hub-test/src/repository/payload"
	"gitlab.com/wit-id/service-hub-test/toolkit/config"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
)

func AddRouteAuthToken(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewAuthTokenService(s.GetDB(), cfg)

	mddw := middleware.NewEnsureToken(s.GetDB(), cfg)

	token := e.Group("/token")
	token.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Auth Token ok")
	})

	token.POST("/auth", authToken(svc))
	token.GET("/refresh", refreshToken(svc), mddw.ValidateRefreshToken)
}

func authToken(svc *service.AuthTokenService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.AuthTokenPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		//Validate request
		if err := request.Validate(); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "validation failed")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		data, err := svc.AuthToken(ctx.Request().Context(), request)
		if err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to auth token")
			return errors.WithStack(httpservice.ErrInternalServerError)
		}

		return httpservice.ResponseData(ctx, payload.ToPayloadAuthToken(data), nil)
	}
}

func refreshToken(svc *service.AuthTokenService) echo.HandlerFunc{
	return func(ctx echo.Context) error {
		data, err := svc.RefreshToken(ctx.Request().Context(), ctx.Get("token-data").(jwt.RequestJWTToken))
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, payload.ToPayloadAuthToken(data), nil)
	}
}