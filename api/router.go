package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, firebaseAuth any) *gin.Engine {
	fmt.Println("Setting up router...")
	r := gin.Default()
	gin.SetMode(os.Getenv("GIN_MODE"))
	r.SetTrustedProxies([]string{"192.168.1.2"})

	//! set db & firebase auth to gin context with a middleware to all incoming requests
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("firebaseAuth", firebaseAuth)
	})

	//setup handlers
	setUpHandlers(r)

	return r
}
