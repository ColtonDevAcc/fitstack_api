package delivery

import (
	"net/http"
	"time"

	"firebase.google.com/go/v4/auth"
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

	client := c.MustGet("FIREBASE_ID_TOKEN").(*auth.Token)
	req.FromUser = client.UID
	req.SentTime = time.Now()

	friendship, err := f.FriendShipUsecase.AddFriend(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, friendship)
}

func (f *FriendshipHandler) RemoveFriend(c *gin.Context) {}

func (f *FriendshipHandler) GetFriends(c *gin.Context) {
	token := c.MustGet("token").(string)

	friends, err := f.FriendShipUsecase.GetFriends(c, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: "err: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, friends)
}
