package controller

import (
	"github.com/VooDooStack/FitStackAPI/domain/model"
	"github.com/gin-gonic/gin"

	"github.com/VooDooStack/FitStackAPI/usecase/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c gin.Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(c gin.Context) error {
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}

	c.JSON(200, u)

	return nil
}
