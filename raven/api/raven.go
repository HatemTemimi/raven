package api

import (
	"fmt"
	raven "github.com/HatemTemimi/Raven/raven/lib"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Raven raven.Raven
}

func (api *Api) Init() {
	fmt.Println("starting server!")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
