package api

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

// https://github.com/VooDooStack/FitStackAPI/blob/dev/router/router.go

func setUpHandlers(router *gin.Engine, v *gin.RouterGroup) {
	v.GET("/ping", func(c *gin.Context) {
		firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)

		c.JSON(200, gin.H{
			"message": "pong",
			"auth":    firebaseAuth,
		})
	}) //.Use(middleware.AuthMiddleware)

	v.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "home",
		})
	})
}
