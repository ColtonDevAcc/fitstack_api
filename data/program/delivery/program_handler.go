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
	g.DELETE("/delete", handler.deleteProgram)
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
	var program program.Program
	client := c.MustGet("FIREBASE_ID_TOKEN").(*auth.Token)
	program.CreatorID = client.UID

	err := c.Bind(&program)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
	}

	err = h.UUsecase.Create(&program)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
	}

	c.JSON(http.StatusOK, program)
}

func (h *ProgramHandler) updateProgram(c *gin.Context) {
	var program program.Program
	c.Bind(&program)
	program.CreatorID = c.GetString("uuid")

	err := h.UUsecase.Update(&program)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
	}

	c.JSON(http.StatusOK, nil)
}

func (h *ProgramHandler) deleteProgram(c *gin.Context) {
	type DeleteProgram struct {
		Id uint `json:"id"`
	}
	var deleteProgram DeleteProgram

	err := c.Bind(&deleteProgram)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
	}

	err = h.UUsecase.Delete(deleteProgram.Id, c.GetString("uuid"))
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}
