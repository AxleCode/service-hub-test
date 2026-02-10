package httpservice

import (
	"net/http"

	"gitlab.com/wit-id/service-hub-test/common/constants"
	"gitlab.com/wit-id/service-hub-test/toolkit/config"

	"github.com/labstack/echo/v4"
)

var Code = "00"
var CodeError = "10"
var Status = "success"
var StatusError = "error"
var MessageEN = "Success"
var MessageID = "Sukses"

type ResponseWithHeader struct {
	AppName  string      `json:"app_name"`
	Version  string      `json:"version"`
	Build    string      `json:"build"`
	Response interface{} `json:"response"`
}

var Message = constants.ErrorResponse{
	ID: "Sukses",
	EN: "Success",
}

type Response struct {
	Data        interface{}             `json:"data"`
	CurrentPage int                     `json:"current_page,omitempty"`
	Limit       int                     `json:"limit,omitempty"`
	TotalPage   int                     `json:"total_page,omitempty"`
	TotalData   int64                   `json:"total_data,omitempty"`
	Message     constants.ErrorResponse `json:"message"`
}

type ResponseCRM struct {
	Code        string      `json:"code"`
	Status      string      `json:"status"`
	Data        interface{} `json:"data"`
	CurrentPage int         `json:"current_page,omitempty"`
	Limit       int         `json:"limit,omitempty"`
	TotalPage   int         `json:"total_page,omitempty"`
	TotalData   int64       `json:"total_data,omitempty"`
	MessageEN   string      `json:"message_en"`
	MessageID   string      `json:"message_id"`
}

//	if err != nil {
//		Message = constants.ErrorResponse{
//			ID: err.Error(),
//			EN: err.Error(),
//		}
//	}
func ResponseData(ctx echo.Context, data interface{}, err error) error {
	return ctx.JSONPretty(http.StatusOK, Response{
		Data:    data,
		Message: Message,
	}, "")
}

func ResponsePagination(ctx echo.Context, data interface{}, err error, page int, limit int, totaPage int, totalData int) error {
	if err != nil {
		Message = constants.ErrorResponse{
			ID: err.Error(),
			EN: err.Error(),
		}
	}

	return ctx.JSONPretty(http.StatusOK, Response{
		Data:        data,
		CurrentPage: page,
		Limit:       limit,
		TotalPage:   totaPage,
		TotalData:   int64(totalData),
		Message:     Message,
	}, "")
}

func ResponseFromJSON(ctx echo.Context, cfg config.KVStore, data interface{}, err error) error {
	if err != nil {
		Code = CodeError
		Status = StatusError
		MessageEN = err.Error()
		MessageID = ""
	}

	return ctx.JSONPretty(http.StatusOK, ResponseWithHeader{
		AppName:  cfg.GetString("app-info.app-name"),
		Version:  cfg.GetString("app-info.version"),
		Build:    cfg.GetString("app-info.build"),
		Response: data,
	}, "")
}

func ResponsePaginationCRM(ctx echo.Context, cfg config.KVStore, data interface{}, err error, page int, limit int, totalPage int, totalData int) error {
	if err != nil {
		Code = CodeError
		Status = StatusError
		MessageEN = err.Error()
		MessageID = ""
	}

	return ctx.JSONPretty(http.StatusOK, ResponseWithHeader{
		AppName: cfg.GetString("app-info.app-name"),
		Version: cfg.GetString("app-info.version"),
		Build:   cfg.GetString("app-info.build"),
		Response: ResponseCRM{
			Code:        Code,
			Status:      Status,
			Data:        data,
			CurrentPage: page,
			Limit:       limit,
			TotalPage:   totalPage,
			TotalData:   int64(totalData),
			MessageEN:   MessageEN,
			MessageID:   MessageID,
		},
	}, "")
}

func ResponseError(ctx echo.Context, cfg config.KVStore, statusCode int, message interface{}) error {
	return ctx.JSONPretty(statusCode, ResponseWithHeader{
		AppName:  cfg.GetString("app-info.app-name"),
		Version:  cfg.GetString("app-info.version"),
		Build:    cfg.GetString("app-info.build"),
		Response: message,
	}, "")
}
