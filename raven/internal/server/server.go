package server

import (
	"github.com/HatemTemimi/Raven/raven/internal/server/routes"
	"github.com/HatemTemimi/Raven/raven/pkg/lib"
	"github.com/labstack/echo/v4"
)

type Api struct {
	raven lib.Raven
}

func (api *Api) Init() {
	raven := api.raven
	raven.Init()
	e := echo.New()
	router := routes.Router{}
	router.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
