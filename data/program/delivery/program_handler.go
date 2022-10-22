package delivery

import (
	"firebase.google.com/go/v4/auth"
	"github.com/VooDooStack/FitStackAPI/api/middleware"
	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type ProgramHandler struct {
	UUsecase domain.ProgramUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewProgramHandler(g *gin.RouterGroup, puc domain.ProgramUsecase, client *auth.Client) {
	handler := &ProgramHandler{
		UUsecase: puc,
	}
	g.GET("/get", middleware.AuthJWT(client), handler.getPrograms)

}

func (h *ProgramHandler) getPrograms(c *gin.Context) {

}
