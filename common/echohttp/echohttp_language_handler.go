package echohttp

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wit-id/service-hub-test/common/constants"
)

func handleLanguage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		lang := c.Request().Header.Get("Accept-Language")
		if lang != constants.LanguageEN && lang != constants.LanguageID {
			c.Request().Header.Set("Accept-Language", constants.LanguageEN)
		}
		return next(c)
	}
}