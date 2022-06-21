package presenter

import "github.com/VooDooStack/FitStackAPI/domain/model"

type CommentPresenter interface {
	ResponseComments(u []*model.Comment) []*model.Comment
}
