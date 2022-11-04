package api

import (
	"fmt"
	"log"
	"os"

	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"
	"github.com/VooDooStack/FitStackAPI/api/middleware"
	"github.com/VooDooStack/FitStackAPI/config"
	_friendshipHandler "github.com/VooDooStack/FitStackAPI/data/friendship/delivery"
	_friendshipRepo "github.com/VooDooStack/FitStackAPI/data/friendship/repository"
	_friendshipUsecase "github.com/VooDooStack/FitStackAPI/data/friendship/usecase"
	_userHandler "github.com/VooDooStack/FitStackAPI/data/user/delivery"
	_userRepo "github.com/VooDooStack/FitStackAPI/data/user/repository"
	_userUseCase "github.com/VooDooStack/FitStackAPI/data/user/usecase"
	"gorm.io/gorm"

	_programHandler "github.com/VooDooStack/FitStackAPI/data/program/delivery"
	_programRepo "github.com/VooDooStack/FitStackAPI/data/program/repository"
	_programUseCase "github.com/VooDooStack/FitStackAPI/data/program/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	client, storage, err := config.SetupFirebase()
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

	setUpHandlers(r, db, *client, storage)

	return r
}

func setUpHandlers(r *gin.Engine, db *gorm.DB, fa auth.Client, storage *storage.Client) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}) //.Use(middleware.AuthMiddleware)

	//===========================Friendship===========================//
	friendshipRG := r.Group("/friendship")
	friendshipRG.Use(middleware.AuthJWT(&fa))
	friendshipRepo := _friendshipRepo.NewFriendshipRepository(*db)
	friendshipUsecase := _friendshipUsecase.NewFriendshipUsecase(friendshipRepo, &fa)
	_friendshipHandler.NewFriendshipHandler(friendshipRG, friendshipUsecase)
	//===========================Friendship===========================//

	//===========================User===========================//
	userRG := r.Group("/user")
	userRepo := _userRepo.NewUserRepository(*db, friendshipRepo)
	userUsecase := _userUseCase.NewUserUseCase(userRepo, fa, storage)
	_userHandler.NewUserHandler(userRG, userUsecase, &fa)
	//===========================User===========================//

	//===========================Program===========================//
	programRG := r.Group("/program")
	programRG.Use(middleware.AuthJWT(&fa))
	programRepo := _programRepo.NewProgramRepository(*db)
	programUsecase := _programUseCase.NewProgramUseCase(programRepo, fa, storage)
	_programHandler.NewProgramHandler(programRG, programUsecase, &fa)
	//===========================Program===========================//
}
