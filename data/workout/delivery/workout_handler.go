package delivery

import (
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/VooDooStack/FitStackAPI/domain/exercise"
	"github.com/gin-gonic/gin"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type WorkoutHandler struct {
	UUsecase exercise.WorkoutUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewWorkoutHandler(g *gin.RouterGroup, wu exercise.WorkoutUsecase, client *auth.Client) {
	handler := &WorkoutHandler{
		UUsecase: wu,
	}

	g.GET("/get", handler.getWorkouts)
	g.POST("/create", handler.createWorkout)
	g.GET("/get-workout", handler.getWorkout)
	g.POST("/update", handler.updateWorkout)
	g.POST("/delete", handler.deleteWorkout)
}

func (wr *WorkoutHandler) getWorkouts(c *gin.Context) {
	workouts, err := wr.UUsecase.GetAll(c.GetString("uuid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, workouts)
}

func (wr *WorkoutHandler) createWorkout(c *gin.Context) {
	workout := exercise.Workout{}
	workout.CreatorId = c.GetString("uuid")

	err := c.ShouldBindJSON(&workout)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})

		return
	}

	creatorId := c.GetString("uuid")
	for i, ws := range workout.WorkoutSets {
		ws.Exercises[i].CreatorID = creatorId
	}

	err = wr.UUsecase.CreateWorkout(&workout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, workout)
}

func (wr *WorkoutHandler) getWorkout(c *gin.Context) {
	workout, err := wr.UUsecase.GetById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, workout)
}

func (wr *WorkoutHandler) updateWorkout(c *gin.Context) {
	var workout exercise.Workout
	err := c.BindJSON(&workout)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})

		return
	}

	err = wr.UUsecase.UpdateWorkout(&workout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, workout)
}

func (wr *WorkoutHandler) deleteWorkout(c *gin.Context) {
	err := wr.UUsecase.DeleteWorkout(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workout deleted"})
}
