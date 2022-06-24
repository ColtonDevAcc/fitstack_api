package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		fmt.Print("Healthy")
	})

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	return r
}
