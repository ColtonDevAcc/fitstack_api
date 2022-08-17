package delivery

import (
	"net/http"

	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	g.GET("/get", handler.FetchUser)
	g.POST("/signup", handler.SignUp)
	g.DELETE("/delete", handler.DeleteUser)
	g.GET("/signin", handler.SignInWithToken)
}

func (ur *UserHandler) FetchUser(c *gin.Context) {
	user, err := ur.UUsecase.GetByUuid(c.Param("uuid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, user)
}

func (ur *UserHandler) SignUp(c *gin.Context) {
	requestedUser := domain.User{}
	err := c.ShouldBindJSON(&requestedUser)
	if err != nil {
		logrus.Error(err)

		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	user, err := ur.UUsecase.SignUp(requestedUser, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})

		return
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

func (ur *UserHandler) SignInWithToken(c *gin.Context) {
	var token struct {
		Token string `json:"token"`
	}

	err := c.ShouldBindJSON(&token)
	if err != nil {
		logrus.Error(err)

		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	user, err := ur.UUsecase.SignInWithToken(c, token.Token)
	if err != nil {
		logrus.Error(err)

		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
