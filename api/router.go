package api

import (
	"fmt"
	"log"
	"os"

	"firebase.google.com/go/v4/auth"
	"github.com/VooDooStack/FitStackAPI/api/middleware"
	"github.com/VooDooStack/FitStackAPI/config"
	_friendshipHandler "github.com/VooDooStack/FitStackAPI/data/friendship/delivery"
	_friendshipRepo "github.com/VooDooStack/FitStackAPI/data/friendship/repository"
	_friendshipUsecase "github.com/VooDooStack/FitStackAPI/data/friendship/usecase"
	_userHandler "github.com/VooDooStack/FitStackAPI/data/user/delivery"
	_userRepo "github.com/VooDooStack/FitStackAPI/data/user/repository"
	_userUseCase "github.com/VooDooStack/FitStackAPI/data/user/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewRouter(db *pgx.Conn) *gin.Engine {
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

	setUpHandlers(r, db, *client)

	return r
}

func setUpHandlers(r *gin.Engine, db *pgx.Conn, fa auth.Client) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}) //.Use(middleware.AuthMiddleware)

	//===========================User===========================//
	userRG := r.Group("/user")
	userRepo := _userRepo.NewUserRepository(*db)
	userUsecase := _userUseCase.NewUserUseCase(userRepo, fa)
	_userHandler.NewUserHandler(userRG, userUsecase, &fa)
	//===========================User===========================//

	//===========================Friendship===========================//
	friendshipRG := r.Group("/friendship")
	friendshipRG.Use(middleware.AuthJWT(&fa))
	friendshipRepo := _friendshipRepo.NewFriendshipRepository(*db)
	friendshipUsecase := _friendshipUsecase.NewFriendshipUsecase(friendshipRepo, &fa)
	_friendshipHandler.NewFriendshipHandler(friendshipRG, friendshipUsecase)
	//===========================Friendship===========================//
}
