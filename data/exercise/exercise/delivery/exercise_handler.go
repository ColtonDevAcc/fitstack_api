package delivery

import (
	"github.com/VooDooStack/FitStackAPI/domain/exercise"
	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ExerciseHandler struct {
	ExerciseShipUsecase exercise.ExerciseUsecase
}

func NewExerciseHandler(g *gin.RouterGroup, us exercise.ExerciseUsecase) {
	handler := &ExerciseHandler{
		ExerciseShipUsecase: us,
	}
	g.GET("/get", handler.GetExercises)
}

func (f *ExerciseHandler) GetExercises(c *gin.Context) {

}
