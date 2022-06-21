package presenter

//! Presenter is a set of a specific implementation of Output Port in Use Cases.
//! Output Port is in charge of handling data from Use cases to the outer layer and defined as abstract.

import "github.com/VooDooStack/FitStackAPI/domain/model"

type commentPresenter struct {
}

type CommentPresenter interface {
	ResponseComments(us []*model.Comment) []*model.Comment
}

func NewCommentPresenter() UserPresenter {
	return &commentPresenter{}
}

func (up *commentPresenter) ResponseUsers(us []*model.Comment) []*model.Comment {
	for _, u := range us {
		u.FirstName = "Mr." + u.FirstName
	}

	return us
}
