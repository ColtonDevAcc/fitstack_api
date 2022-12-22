package delivery

import (
	"net/http"

	"github.com/VooDooStack/FitStackAPI/domain/muscle"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type RecoveryHandler struct {
	RUsecase muscle.RecoveryUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewRecoveryHandler(g *gin.RouterGroup, ru muscle.RecoveryUsecase) {
	handler := &RecoveryHandler{
		RUsecase: ru,
	}
	g.GET("/stats", handler.FetchRecovery)

}

func (ur *RecoveryHandler) FetchRecovery(c *gin.Context) {
	recovery, err := ur.RUsecase.FetchRecovery(c.GetString("uuid"))
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, recovery)
}
