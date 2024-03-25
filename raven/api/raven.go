package api

import (
	"fmt"
	"github.com/HatemTemimi/Raven/raven/lib"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Raven lib.Raven
}

func (api *Api) Init() {
	fmt.Println("starting server!")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "raven api",
		})
	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
