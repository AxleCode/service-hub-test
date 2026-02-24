package application

import (
	"database/sql"
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/constants"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/common/jwt"
	"gitlab.com/wit-id/service-hub-test/common/utility"
	"gitlab.com/wit-id/service-hub-test/src/authentication/service"
	"gitlab.com/wit-id/service-hub-test/src/middleware"
	"gitlab.com/wit-id/service-hub-test/src/repository/payload"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/service-hub-test/toolkit/config"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
	// smtp "gitlab.com/wit-id/service-hub-test/toolkit/smtp"
)

func AddRouteAuthentication(s *httpservice.Service, cfg config.KVStore, e *echo.Echo){
	svc := service.NewAuthenticationService(s.GetDB(), cfg)
	mddw := middleware.NewEnsureToken(s.GetDB(), cfg)

	authApp := e.Group("/auth")
	authApp.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "auth app ok")
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

	authApp.Use(mddw.ValidateToken)

	authApp.POST("/login", login(svc))
	authApp.POST("/register", register(svc))
	authApp.POST("/logout", logout(svc))
}

func register(svc *service.AuthenticationService) echo.HandlerFunc{
	return func(ctx echo.Context) error {
		var (
			request			payload.RegisterPayload
			result			sqlc.GetAuthenticationByIDRow
			
		)
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(ctx.Request().Context()); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		customerData, err := svc.UserRegisterFromWeb(ctx.Request().Context(), request.ToEntity())
		if err != nil {
			return err
		}

		salt, _ := utility.GenerateSaltLaravelBase64()
		storedSalt, _ := base64.StdEncoding.DecodeString(salt)

		authRequest := sqlc.InsertAuthenticationParams{
			Guid: 		utility.GenerateGoogleUUID(),
			UserGuid: 	sql.NullString{String: customerData.Guid, Valid: true},
			Username: 	request.Username,
			Salt:		sql.NullString{String: salt, Valid: true},
			Status:		constants.StatusActive,
			Password: 	utility.HashPasswordLaravel(request.Password, storedSalt),
			CreatedBy: 	constants.CreatedByTemporaryBySystem,	
		}

		// log.FromCtx(ctx.Request().Context()).Info("hash password", authRequest.Password)
		
		authResult, err := svc.AuthenticationRegister(ctx.Request().Context(), authRequest)
		if err != nil{
			return err
		}

		result, err = svc.ReadAuthenticationByID(ctx.Request().Context(), authResult.Guid)
		if err != nil{
			return err
		}

		return httpservice.ResponseData(ctx, payload.ToPayloadUserProfile(result), nil)
		// return httpservice.ResponseData(ctx, authResult, nil)
	}
}

func login(svc *service.AuthenticationService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.LoginPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(ctx.Request().Context()); err != nil {
			return err
		}

		data, err := svc.Login(ctx.Request().Context(), request, ctx.Get("token-data").(jwt.RequestJWTToken))
		if err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return err
		}

		return httpservice.ResponseData(ctx, payload.ToPayloadAuthentication(data), nil)
	}
}

func logout(svc *service.AuthenticationService) echo.HandlerFunc{
	return func(ctx echo.Context) error {
		err := svc.LogoutToken(ctx.Request().Context(), ctx.Get("token-data").(jwt.RequestJWTToken))
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, nil, nil)
	}
}