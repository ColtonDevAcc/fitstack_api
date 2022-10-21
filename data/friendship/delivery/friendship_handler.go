package delivery

import (
	"fmt"
	"net/http"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	g.POST("/get-friend", handler.GetFriend)
	g.GET("/get-friends-list", handler.GetFriendsList)
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

	err := f.FriendShipUsecase.AddFriend(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (f *FriendshipHandler) RemoveFriend(c *gin.Context) {}

func (f *FriendshipHandler) GetFriends(c *gin.Context) {
	client := c.MustGet("FIREBASE_ID_TOKEN").(*auth.Token)

	friends, err := f.FriendShipUsecase.GetFriends(c, client.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: "err: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, friends)
}

func (f *FriendshipHandler) GetFriendsList(c *gin.Context) {
	client := c.MustGet("FIREBASE_ID_TOKEN").(*auth.Token)

	friends, err := f.FriendShipUsecase.GetFriendsList(c, client.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: "err: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, friends)
}

func (f *FriendshipHandler) GetFriend(c *gin.Context) {
	type emailStruct struct {
		Email string `json:"email"`
	}
	email := emailStruct{}
	err := c.ShouldBindJSON(&email)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, ResponseError{Message: "err: " + err.Error()})
		return
	}

	fmt.Printf("here is the email %v", email.Email)

	friends, err := f.FriendShipUsecase.GetFriend(c, email.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: "err: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, friends)
}
