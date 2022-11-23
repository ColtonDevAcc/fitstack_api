package delivery

import (
	"fmt"

	"firebase.google.com/go/v4/auth"
	"github.com/VooDooStack/FitStackAPI/domain/exercise"
	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ExerciseHandler struct {
	ExerciseUsecase exercise.ExerciseUsecase
}

func NewExerciseHandler(g *gin.RouterGroup, us exercise.ExerciseUsecase) {
	handler := &ExerciseHandler{
		ExerciseUsecase: us,
	}
	g.GET("/get", handler.GetExercises)
	g.POST("/create", handler.CreateExercise)
	g.POST("/update", handler.UpdateExercise)
	g.POST("/delete", handler.DeleteExercise)
	g.GET("/", handler.GetExercise)
}

func (f *ExerciseHandler) GetExercises(c *gin.Context) {
	exercises, err := f.ExerciseUsecase.GetExercises(c.GetString("uuid"))
	if err != nil {
		c.JSON(500, ResponseError{Message: err.Error()})

		return
	}

	c.JSON(200, exercises)
}

func (f *ExerciseHandler) CreateExercise(c *gin.Context) {
	exercise := exercise.Exercise{}

	err := c.ShouldBindJSON(&exercise)
	if err != nil {
		c.JSON(400, ResponseError{Message: err.Error()})

		return
	}

	exercise.CreatorID = c.GetString("uuid")

	err = f.ExerciseUsecase.CreateExercise(&exercise)
	if err != nil {
		c.JSON(500, ResponseError{Message: err.Error()})

		return
	}

	c.JSON(200, exercise)
}

func (f *ExerciseHandler) UpdateExercise(c *gin.Context) {
	exercise := exercise.Exercise{}
	client := c.MustGet("FIREBASE_ID_TOKEN").(*auth.Token)

	err := c.ShouldBindJSON(&exercise)
	if err != nil {
		c.JSON(400, ResponseError{Message: err.Error()})

		return
	}

	fmt.Printf("\nexercise: %+v \n", exercise)

	if exercise.CreatorID != client.UID {
		c.JSON(403, ResponseError{Message: "You are not the creator of this exercise"})
		return
	}

	err = f.ExerciseUsecase.UpdateExercise(&exercise)
	if err != nil {
		c.JSON(500, ResponseError{Message: err.Error()})

		return
	}

	c.JSON(200, exercise)
}

func (f *ExerciseHandler) DeleteExercise(c *gin.Context) {
	exercise := exercise.Exercise{}
	exercise.CreatorID = c.GetString("uuid")

	err := c.ShouldBindJSON(&exercise)
	if err != nil {
		c.JSON(400, ResponseError{Message: err.Error()})

		return
	}

	err = f.ExerciseUsecase.DeleteExercise(&exercise)
	if err != nil {
		c.JSON(500, ResponseError{Message: err.Error()})

		return
	}

	c.JSON(200, exercise)
}

func (f *ExerciseHandler) GetExercise(c *gin.Context) {
	exercise, err := f.ExerciseUsecase.GetExercises(c.GetString("uuid"))
	if err != nil {
		c.JSON(500, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(200, exercise)
}
