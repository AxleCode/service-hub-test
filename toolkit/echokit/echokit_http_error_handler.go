package echokit

import (
	"github.com/labstack/echo/v4"
)

func loggerHTTPErrorHandler(next echo.HTTPErrorHandler) echo.HTTPErrorHandler {
	return next
}