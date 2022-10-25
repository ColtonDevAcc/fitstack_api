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
