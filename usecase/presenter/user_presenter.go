package presenter

import "github.com/VooDooStack/FitStackAPI/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
