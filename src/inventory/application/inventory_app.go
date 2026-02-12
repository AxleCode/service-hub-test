package application

import (
	"net/http"
	"net/url"
	"math"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/src/inventory/service"
	"gitlab.com/wit-id/service-hub-test/src/repository/payload"
	"gitlab.com/wit-id/service-hub-test/toolkit/config"
	"github.com/labstack/echo/v4"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
)

func AddRouteInventory(s *httpservice.Service, cfg config.KVStore, e *echo.Echo){
	svc := service.NewInventoryService(s.GetDB(), cfg)

	inventoryApp := e.Group("/inventory")

	inventoryApp.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Inventory ok")
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

	inventoryApp.POST("", createInventory(svc, client, cfg))
	inventoryApp.GET("/detail/:guid", getInventory(svc, client, cfg))
	inventoryApp.POST("/list", listInventory(svc, client, cfg))
}

func createInventory(svc *service.InventoryService, client http.Client, cfg config.KVStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request payload.InsertInventoryPayload
		if err := c.Bind(&request); err != nil {
			log.FromCtx(c.Request().Context()).Error(err, "failed parse request body")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		inventoryData, err := svc.CreateInventory(c.Request().Context(), request.ToEntity(), request)
		if err != nil {
			return err
		}

		data, err := svc.GetInventory(c.Request().Context(), inventoryData)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(c, data, nil)
	}
}

func getInventory(svc *service.InventoryService, client http.Client, cfg config.KVStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		guid := c.Param("guid")

		data,  err := svc.GetInventory(c.Request().Context(), guid)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(c, data, nil)
	}
}

func listInventory(svc *service.InventoryService, client http.Client, cfg config.KVStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request payload.ListInventoryPayload
		if err := c.Bind(&request); err != nil {
			log.FromCtx(c.Request().Context()).Error(err, "failed parse request body")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(); err != nil {
			return err 
		}
		
		inventoryData, totalData, err := svc.ListInventory(c.Request().Context(), request.ToEntity())
		if err != nil {
			log.FromCtx(c.Request().Context()).Error(err, "failed get list inventory")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		//TOTAL PAGE
		totalPage := math.Ceil(float64(totalData) / float64(request.Limit))

		return httpservice.ResponsePagination(c,
			payload.ToPayloadListInventory(inventoryData),
			nil, int(request.Offset),
			int(request.Limit),
			int(totalPage),
			int(totalData))
	}
}