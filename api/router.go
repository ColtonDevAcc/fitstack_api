package api

import (
	"fmt"
	"log"
	"os"

	"github.com/VooDooStack/FitStackAPI/config"
	_userHandler "github.com/VooDooStack/FitStackAPI/data/user/delivery"
	_userRepo "github.com/VooDooStack/FitStackAPI/data/user/repository"
	_userUseCase "github.com/VooDooStack/FitStackAPI/data/user/usecase"
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

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("firebaseAuth", client)
	})

	gin.SetMode(os.Getenv("GIN_MODE"))
	r.SetTrustedProxies([]string{"192.168.1.2"})

	setUpHandlers(r, db)

	return r
}

func setUpHandlers(r *gin.Engine, db *gorm.DB) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}) //.Use(middleware.AuthMiddleware)

	//===========================User===========================//
	userRG := r.Group("/user")
	userRepo := _userRepo.NewUserRepository(*db)
	userUsecase := _userUseCase.NewUserUseCase(userRepo)
	_userHandler.NewUserHandler(userRG, userUsecase)
	//===========================User===========================//
}
