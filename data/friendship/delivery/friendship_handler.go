package delivery

import (
	"net/http"

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
}

func (f *FriendshipHandler) AddFriend(c *gin.Context) {
	var req domain.Friendship
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, ResponseError{Message: "Invalid request"})
		return
	}

	friendship, err := f.FriendShipUsecase.AddFriend(req)
	if err != nil {
		c.JSON(400, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, friendship)
}

func (f *FriendshipHandler) RemoveFriend(c *gin.Context) {}
