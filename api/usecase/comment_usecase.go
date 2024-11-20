package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type ICommentUsecase interface {
	CreateComment(comment model.Comment) (model.CommentResponse, error)
}

type commentUsecase struct {
	cr repository.ICommentRepository
	cv validator.ICommentValidator
}

func NewCommentUsecase(cr repository.ICommentRepository, cv validator.ICommentValidator) ICommentUsecase {
	return &commentUsecase{cr, cv}
}

func (cu *commentUsecase) CreateComment(comment model.Comment) (model.CommentResponse, error) {
	if err := cu.cv.CommentValidate(comment); err != nil {
		return model.CommentResponse{}, err
	}

	if err := cu.cr.CreateComment(&comment); err != nil {
		return model.CommentResponse{}, err
	}

	resComment := model.CommentResponse{
		ID:        comment.ID,
		Comment:   comment.Comment,
		UserId:    comment.UserId,
		TweetId:   comment.TweetId,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
	return resComment, nil
}
