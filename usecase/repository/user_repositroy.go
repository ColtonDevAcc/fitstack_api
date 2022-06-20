package repository

import "github.com/VooDooStack/FitStackAPI/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
