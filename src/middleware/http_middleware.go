package middleware

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/wit-id/service-hub-test/common/constants"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/common/jwt"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/service-hub-test/toolkit/config"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
)

type EnsureToken struct {
	mainDB *sql.DB
	config config.KVStore
}

type AccessPage struct {
	Page    string   `json:"page"`
	KeyPage string   `json:"key_page"`
	Access  []string `json:"access"`
}

//var allLanguage bool

func NewEnsureToken(db *sql.DB, cfg config.KVStore) *EnsureToken {
	return &EnsureToken{
		mainDB: db,
		config: cfg,
	}
}

func (v *EnsureToken) ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := ctx.Request()

		headerDataToken := request.Header.Get(v.config.GetString("header.token-param"))
		if headerDataToken == "" {
			// return echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgHeaderTokenNotFound).SetInternal(errors.Wrap(httpservice.ErrMissingHeaderData, httpservice.MsgHeaderTokenNotFound))
			return ctx.JSON(http.StatusUnauthorized, echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgHeaderTokenNotFound))
		}

		jwtResponse, err := jwt.ClaimsJwtToken(ctx.Request().Context(), v.config, headerDataToken)
		if err != nil {
			// return echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgHeaderTokenUnauthorized).SetInternal(errors.Wrap(err, httpservice.MsgHeaderTokenUnauthorized))
			return ctx.JSON(http.StatusUnauthorized, echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgHeaderTokenUnauthorized))
		}

		// Set data jwt response to ...
		ctx.Set(constants.MddwTokenKey, jwtResponse)

		return next(ctx)
	}
}

func (v *EnsureToken) ValidateRefreshToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := ctx.Request()

		headerDataToken := request.Header.Get(v.config.GetString("header.refresh-token-param"))
		if headerDataToken == "" {
			// return echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgHeaderRefreshTokenNotFound).SetInternal(errors.Wrap(httpservice.ErrMissingHeaderData, httpservice.MsgHeaderRefreshTokenNotFound))
			return ctx.JSON(http.StatusUnauthorized, echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgHeaderRefreshTokenNotFound))
		}

		jwtResponse, err := jwt.ClaimsJwtToken(ctx.Request().Context(), v.config, headerDataToken)
		if err != nil {
			// return echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgHeaderRefreshTokenUnauthorized).SetInternal(errors.Wrap(err, httpservice.MsgHeaderRefreshTokenUnauthorized))
			return ctx.JSON(http.StatusUnauthorized, echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgHeaderRefreshTokenUnauthorized))
		}

		// Set data jwt response to ...
		ctx.Set(constants.MddwTokenKey, jwtResponse)

		return next(ctx)
	}
}

func (v *EnsureToken) ValidateUserLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Get data token session
		tokenAuth := ctx.Get(constants.MddwTokenKey).(jwt.RequestJWTToken)

		q := sqlc.New(v.mainDB)

		tokenData, err := q.GetAuthToken(ctx.Request().Context(), sqlc.GetAuthTokenParams{
			Name:       tokenAuth.AppName,
			DeviceID:   tokenAuth.DeviceID,
			DeviceType: tokenAuth.DeviceType,
		})
		if err != nil {
			// return echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgHeaderTokenUnauthorized).SetInternal(errors.Wrap(httpservice.ErrUnauthorizedTokenData, httpservice.MsgHeaderTokenUnauthorized))
			return ctx.JSON(http.StatusUnauthorized, echo.NewHTTPError(http.StatusUnauthorized, httpservice.ErrMissingHeaderData))
		}

		if !tokenData.IsLogin {
			// return echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgIsNotLogin).SetInternal(errors.WithMessage(httpservice.ErrUnauthorizedUser, httpservice.MsgIsNotLogin))
			return ctx.JSON(http.StatusUnauthorized, echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgIsNotLogin))
		}

		// Get user authentication
		userData, err := q.GetAuthenticationIAMByID(ctx.Request().Context(), tokenData.UserLogin.String)
		if err != nil {
			//  return echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgUnauthorizedUser).SetInternal(errors.Wrap(httpservice.ErrUnauthorizedUser, httpservice.MsgUnauthorizedUser))
			log.FromCtx(ctx.Request().Context()).Error(err, "failed get authentication IAM by id")
			return ctx.JSON(http.StatusUnauthorized, echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgUnauthorizedUser))
		}

		// // check active user {
		if userData.Status != "active" {
			// return echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgUserNotActive).SetInternal(errors.WithMessage(httpservice.ErrUnauthorizedUser, httpservice.MsgUserNotActive))
			return ctx.JSON(http.StatusUnauthorized, echo.NewHTTPError(http.StatusUnauthorized, httpservice.MsgUserNotActive))
		}

		// Set data user response to ...
		ctx.Set(constants.MddwUserBackoffice, userData)

		// Set data IAM
		// filterIAM := IAMAccessPayload{
		// 	SiePelayananID:       userData.SiePelayananID.String,
		// 	DepartmentID:         userData.DepartmentID.String,
		// 	DivisiID:             userData.DivisiID.String,
		// 	FamilyAltarAnggotaID: userData.PengurusID.String,
		// }

		// iamData, err := v.getIAMAccessToken(filterIAM)
		// if err == nil {
		// 	if userData.SiePelayananID.Valid {
		// 		property, errProperty := q.GetMasterdataValues(context.Background(), userData.SiePelayananID.String)
		// 		if errProperty == nil {
		// 			iamData.SiePalayananID.ID = property.Guid
		// 			iamData.SiePalayananID.Name = property.Value
		// 		} else {
		// 			log.Println("faield get property", errProperty)
		// 		}

		// 	}

		// 	ctx.Set(constants.MddwKeyRole, iamData)
		// }

		return next(ctx)
	}
}
