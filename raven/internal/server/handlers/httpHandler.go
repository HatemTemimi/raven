package handlers

import (
	"github.com/HatemTemimi/Raven/raven/pkg/lib"
	"github.com/HatemTemimi/Raven/raven/views"
	"github.com/a-h/templ"
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

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func (handler *HttpHandler) Index(c echo.Context) error {
	proxies, _ := handler.Raven.FetchAll()
	return renderView(c, views.Index(proxies))
}
