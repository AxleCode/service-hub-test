package application

import (
	"net/http"
	"net/url"
	"math"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/src/barang/service"
	"gitlab.com/wit-id/service-hub-test/src/repository/payload"
	"gitlab.com/wit-id/service-hub-test/toolkit/config"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
)

func AddRouteBarang(s *httpservice.Service, cfg config.KVStore, e *echo.Echo){
	svc := service.NewBarangService(s.GetDB(), cfg)

	barangApp := e.Group("/barang")

	barangApp.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Barang ok")
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

	barangApp.POST("", createBarang(svc, client, cfg))
	barangApp.GET("/detail/:guid", getBarang(svc, client, cfg))
	barangApp.POST("/list", listBarang(svc))
	barangApp.PUT("/:guid", updateBarang(svc))
	barangApp.DELETE("/:guid", deleteBarang(svc))
}

func createBarang(svc *service.BarangService, client http.Client, cfg config.KVStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request payload.InsertBarangPayload
		if err := c.Bind(&request); err != nil {
			log.FromCtx(c.Request().Context()).Error(err, "failed parse request body")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		barangData, err := svc.CreateBarang(c.Request().Context(), request.ToEntity(), request)
		if err != nil {
			return err
		}

		data, err := svc.GetBarang(c.Request().Context(), barangData)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(c, data, nil)
	}
}

func getBarang(svc *service.BarangService, client http.Client, cfg config.KVStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		guid := c.Param("guid")

		barangData, err := svc.GetBarang(c.Request().Context(), guid)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(c, barangData, nil)
	}
}

func listBarang(svc *service.BarangService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request payload.ListBarangPayload
		if err := c.Bind(&request); err != nil {
			log.FromCtx(c.Request().Context()).Error(err, "failed parse request body")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(); err != nil {
			return err
		}

		listData, totalData, err := svc.ListBarang(c.Request().Context(), request.ToEntity())
		if err != nil {
			return err
		}

		//TOTAL PAGE
		totalPage := math.Ceil(float64(totalData) / float64(request.Limit))

		return httpservice.ResponsePagination(c,
			payload.ToPayloadListBarang(listData),
			nil, int(request.Offset),
			int(request.Limit),
			int(totalPage),
			int(totalData))
	}
}

func updateBarang(svc *service.BarangService) echo.HandlerFunc {
	return func(c echo.Context) error {
		guid := c.Param("guid")
		if guid == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		var request payload.UpdateBarangPayload
		if err := c.Bind(&request); err != nil {
			log.FromCtx(c.Request().Context()).Error(err, "failed parse request body")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(); err != nil {
			return err
		}

		barang, err := svc.UpdateBarang(c.Request().Context(),
			request.ToEntity(guid))
		if err != nil {
			return err
		}

		return httpservice.ResponseData(
			c,
			payload.ToPayloadBarang(barang),
			nil)
	}
}

func deleteBarang(svc *service.BarangService) echo.HandlerFunc {
	return func(c echo.Context) error {
		guid := c.Param("guid")
		if guid == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		err := svc.DeleteBarang(c.Request().Context(), guid)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(
			c,
			nil,
			nil)
	}
}