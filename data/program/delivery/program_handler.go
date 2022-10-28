package delivery

import (
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/VooDooStack/FitStackAPI/domain/program"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type ProgramHandler struct {
	UUsecase program.ProgramUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewProgramHandler(g *gin.RouterGroup, puc program.ProgramUsecase, client *auth.Client) {
	handler := &ProgramHandler{
		UUsecase: puc,
	}
	g.GET("/get", handler.getPrograms)
	g.POST("/create", handler.createProgram)
	g.POST("/update", handler.updateProgram)
}

func (h *ProgramHandler) getPrograms(c *gin.Context) {
	uuid := c.GetString("uuid")
	program, err := h.UUsecase.Get(uuid)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
	}

	c.JSON(http.StatusOK, program)
}

func (h *ProgramHandler) createProgram(c *gin.Context) {
	uuid := c.GetString("uuid")
	var program program.Program
	program.Creator = uuid

	c.Bind(&program)
	err := h.UUsecase.Create(&program)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *ProgramHandler) updateProgram(c *gin.Context) {
	var program program.Program
	c.Bind(&program)

	if program.Creator != c.GetString("uuid") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
			"error":   true,
		})

		return
	}

	err := h.UUsecase.Update(&program)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}
