package echohttp

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"gitlab.com/wit-id/service-hub-test/toolkit/config"
)

func handleEchoError(_ config.KVStore) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			_ = c.JSON(he.Code, he.Message)
			return
		}
		_ = c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error"})
	}
}