package server

import (
	"fmt"
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
	fmt.Println("starting server!")
	e := echo.New()
	router := routes.Router{}
	router.Init(e)
	router.InitRoutes()
	e.Logger.Fatal(e.Start(":8080"))
	/*
		r := gin.Default()
			r.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "raven server",
				})
			})
			r.GET("/get-all", func(c *gin.Context) {
				proxies, err := raven.FetchAll()
				if err != nil {
					c.JSON(500, gin.H{
						"error": err,
					})
				} else {
					c.JSON(200, gin.H{
						"proxies": proxies,
					})
				}

			})
			err := r.Run()
			if err != nil {
				return
			} // listen and serve on 0.0.0.0:8080

	*/
}
