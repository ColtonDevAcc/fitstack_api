package presenter

//! Presenter is a set of a specific implementation of Output Port in Use Cases.
//! Output Port is in charge of handling data from Use cases to the outer layer and defined as abstract.

import "github.com/VooDooStack/FitStackAPI/domain/model"

type userPresenter struct {
}

type UserPresenter interface {
	ResponseUsers(us []*model.User) []*model.User
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.FirstName = "Mr." + u.FirstName
	}

	return us
}