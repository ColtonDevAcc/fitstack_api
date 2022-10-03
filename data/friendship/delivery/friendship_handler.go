package delivery

import (
	"net/http"
	"strings"

	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

type FriendshipHandler struct {
	FriendShipUsecase domain.FriendshipUsecase
}

func NewFriendshipHandler(g *gin.RouterGroup, us domain.FriendshipUsecase) {
	handler := &FriendshipHandler{
		FriendShipUsecase: us,
	}
	g.POST("/add", handler.AddFriend)
	g.POST("/remove", handler.RemoveFriend)
	g.GET("/get-friends", handler.GetFriends)
}

func (f *FriendshipHandler) AddFriend(c *gin.Context) {
	var req domain.Friendship
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, ResponseError{Message: err.Error()})
		return
	}

	friendship, err := f.FriendShipUsecase.AddFriend(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, friendship)
}

func (f *FriendshipHandler) RemoveFriend(c *gin.Context) {}

func (f *FriendshipHandler) GetFriends(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.String(http.StatusForbidden, "No Authorization header provided")
		c.Abort()
		return
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		c.String(http.StatusForbidden, "Could not find bearer token in Authorization header")
		c.Abort()
		return
	}

	friends, err := f.FriendShipUsecase.GetFriends(c, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: "err: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, friends)
}
