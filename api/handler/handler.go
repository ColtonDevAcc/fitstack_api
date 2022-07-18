package handler

import (
	"github.com/gin-gonic/gin"
)

// https://github.com/VooDooStack/FitStackAPI/blob/dev/router/router.go

//! TODO: better var names for groups
func SetUpHandlers(v *gin.RouterGroup) {
	v.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	}) //.Use(middleware.AuthMiddleware)

	v.GET("/signIn", userSignIn)
	v.GET("/signUp", signUp)
}
