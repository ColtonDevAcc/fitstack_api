package repository

import "github.com/VooDooStack/FitStackAPI/domain/model"

type CommentRepository interface {
	FindAll(u []*model.Comment) ([]*model.Comment, error)
}
