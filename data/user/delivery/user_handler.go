package delivery

import (
	"net/http"

	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/gin-gonic/gin"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type UserHandler struct {
	UUsecase domain.UserUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewUserHandler(g *gin.RouterGroup, us domain.UserUsecase) {
	handler := &UserHandler{
		UUsecase: us,
	}
	g.GET("/get/", handler.FetchUser)
	g.POST("/signUp/", handler.SignUp)
	g.DELETE("/delete/", handler.DeleteUser)
}

func (ur *UserHandler) FetchUser(c *gin.Context) {
	user, err := ur.UUsecase.GetByUuid(c.Param("uuid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	c.JSON(http.StatusOK, user)
}

func (ur *UserHandler) SignUp(c *gin.Context) {
	//TODO: implement sign up
	user, err := ur.UUsecase.GetByUuid(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	c.JSON(http.StatusOK, user)
}

func (ur *UserHandler) DeleteUser(c *gin.Context) {
	err := ur.UUsecase.Delete(c.Param("uuid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}
