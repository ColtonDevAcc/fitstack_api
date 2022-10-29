package delivery

import (
	"github.com/VooDooStack/FitStackAPI/domain/exercise"
	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

type WorkoutHandler struct {
	WorkoutUsecase exercise.WorkoutUsecase
}

func NewWorkoutHandler(g *gin.RouterGroup, us exercise.WorkoutUsecase) {
	handler := &WorkoutHandler{
		WorkoutUsecase: us,
	}
	g.GET("/get", handler.Get)
	g.POST("/create", handler.CreateWorkout)
}

func (f *WorkoutHandler) Get(c *gin.Context) {

}

func (f *WorkoutHandler) CreateWorkout(c *gin.Context) {
	var req exercise.Workout
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, ResponseError{Message: err.Error()})
		return
	}

	err := f.WorkoutUsecase.CreateWorkout(&req)
	if err != nil {
		c.JSON(400, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(200, nil)
}
