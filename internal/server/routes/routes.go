package routes

import (
	"github.com/HatemTemimi/raven/internal/server/handlers"
	"github.com/HatemTemimi/raven/pkg/lib"
	Checker "github.com/HatemTemimi/raven/pkg/lib/checker"
	Reader "github.com/HatemTemimi/raven/pkg/lib/reader"
	Scanner "github.com/HatemTemimi/raven/pkg/lib/scanner"
	Writer "github.com/HatemTemimi/raven/pkg/lib/writer"
	"github.com/labstack/echo/v4"
)

type Router struct {
	echo    *echo.Echo
	handler *handlers.HttpHandler
}

func (router *Router) Init(e *echo.Echo) {
	Raven := lib.Raven{
		Scanner: Scanner.Scanner{},
		Checker: Checker.Checker{},
		Writer:  Writer.Writer{},
		Reader:  Reader.Reader{},
	}
	Raven.Init()
	router.handler = &handlers.HttpHandler{
		Raven: Raven,
	}
	router.echo = e
	router.InitRoutes()
}

func (router *Router) InitRoutes() {
	router.echo.GET("api/fetch", router.handler.FetchAll)
	router.echo.GET("api/fetch/:target", router.handler.FetchValid)
}
