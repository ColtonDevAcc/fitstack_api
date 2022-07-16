package api

import (
	"fmt"
	"log"
	"os"

	"github.com/VooDooStack/FitStackAPI/api/middleware"
	"github.com/VooDooStack/FitStackAPI/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	client, err := config.SetupFirebase()
	if err != nil {
		log.Fatalln("failed to init firebase auth", err)
	}

	fmt.Println("Setting up router...")
	r := gin.Default()

	v := r.Group("/v1")
	// authGroup := r.Group("/auth")

	v.Use(middleware.AuthJWT(client))
	gin.SetMode(os.Getenv("GIN_MODE"))
	r.SetTrustedProxies([]string{"192.168.1.2"})

	setUpHandlers(r, v)

	return r
}
