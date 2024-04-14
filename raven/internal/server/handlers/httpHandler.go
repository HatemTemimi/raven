package handlers

import (
	"github.com/HatemTemimi/Raven/raven/pkg/lib"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HttpHandler struct {
	Raven lib.Raven
}

func (handler *HttpHandler) FetchAll(c echo.Context) error {
	proxies, err := handler.Raven.FetchAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, proxies)
}

func (handler *HttpHandler) FetchValid(c echo.Context) error {
	target := c.Param("target")
	proxies, err := handler.Raven.FetchValid(target)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, proxies)
}
