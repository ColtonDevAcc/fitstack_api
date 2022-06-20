package controller

import (
	"net/http"

	"github.com/VooDooStack/FitStackAPI/domain/model"
	"github.com/gorilla/mux"

	"github.com/VooDooStack/FitStackAPI/usecase/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c mux.Router) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(c mux.Router) error {
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}

	return Json(http.StatusOK, u)
}
