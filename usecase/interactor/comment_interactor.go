package interactor

import (
	"github.com/VooDooStack/FitStackAPI/domain/model"
	"github.com/VooDooStack/FitStackAPI/usecase/presenter"
	"github.com/VooDooStack/FitStackAPI/usecase/repository"
)

type commentInteractor struct {
	CommentRepository repository.CommentRepository
	CommentPresenter  presenter.CommentPresenter
}

type CommentInteractor interface {
	GetComments(u []*model.Comment) ([]*model.Comment, error)
}

func NewCommentInteractor(r repository.CommentRepository, p presenter.CommentPresenter) CommentInteractor {
	return &commentInteractor{r, p}
}

func (us *commentInteractor) GetComments(u []*model.Comment) ([]*model.Comment, error) {
	u, err := us.CommentRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.CommentPresenter.ResponseComments(u), nil
}
