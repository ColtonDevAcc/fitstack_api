package api

import (
	"fmt"
	"log"
	"os"

	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"
	"github.com/VooDooStack/FitStackAPI/api/middleware"
	"github.com/VooDooStack/FitStackAPI/config"
	_workoutHandler "github.com/VooDooStack/FitStackAPI/data/exercise/workout/delivery"
	_workoutRepo "github.com/VooDooStack/FitStackAPI/data/exercise/workout/repository"
	_workoutUseCase "github.com/VooDooStack/FitStackAPI/data/exercise/workout/usecase"

	_exerciseHandler "github.com/VooDooStack/FitStackAPI/data/exercise/exercise/delivery"
	_exerciseRepo "github.com/VooDooStack/FitStackAPI/data/exercise/exercise/repository"
	_exerciseUseCase "github.com/VooDooStack/FitStackAPI/data/exercise/exercise/usecase"

	_friendshipHandler "github.com/VooDooStack/FitStackAPI/data/friendship/delivery"
	_friendshipRepo "github.com/VooDooStack/FitStackAPI/data/friendship/repository"
	_friendshipUsecase "github.com/VooDooStack/FitStackAPI/data/friendship/usecase"
	_programHandler "github.com/VooDooStack/FitStackAPI/data/program/delivery"
	_programRepo "github.com/VooDooStack/FitStackAPI/data/program/repository"
	_programUseCase "github.com/VooDooStack/FitStackAPI/data/program/usecase"
	_userHandler "github.com/VooDooStack/FitStackAPI/data/user/delivery"
	_userRepo "github.com/VooDooStack/FitStackAPI/data/user/repository"
	_userUseCase "github.com/VooDooStack/FitStackAPI/data/user/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	//===========================Workout===========================//
	workoutRG := r.Group("/workout")
	workoutRG.Use(middleware.AuthJWT(&fa))
	workoutRepo := _workoutRepo.NewWorkoutRepository(*db)
	workoutUsecase := _workoutUseCase.NewWorkoutUseCase(workoutRepo)
	_workoutHandler.NewWorkoutHandler(workoutRG, workoutUsecase, &fa)
	//===========================Workout===========================//

	//===========================Exercise===========================//
	exerciseRG := r.Group("/exercise")
	exerciseRG.Use(middleware.AuthJWT(&fa))
	exerciseRepo := _exerciseRepo.NewExerciseRepository(*db)
	exerciseUsecase := _exerciseUseCase.NewExerciseUsecase(exerciseRepo)
	_exerciseHandler.NewExerciseHandler(exerciseRG, exerciseUsecase)
	//===========================Exercise===========================//
}
