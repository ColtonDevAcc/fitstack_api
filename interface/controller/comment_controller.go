package controller

//! Controllers are a set of a specific implementation of Input Port in Use Cases.
//! Input Port is in charge of handling data from the outer layer and defined as abstract.

import (
	"github.com/VooDooStack/FitStackAPI/domain/model"
	"github.com/gin-gonic/gin"

	"github.com/VooDooStack/FitStackAPI/usecase/interactor"
)

type commentController struct {
	CommentInteractor interactor.CommentInteractor
}

type CommentController interface {
	GetComment(c gin.Context) error
}

func NewCommentController(us interactor.UserInteractor) CommentController {
	return &commentController{us}
}

func (uc *commentController) GetComment(c gin.Context) error {
	var u []*model.Comment

	u, err := uc.CommentInteractor.GetComments(u)
	if err != nil {
		return err
	}

	c.JSON(200, u)

	return nil
}
