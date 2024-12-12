package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"

	"gorm.io/gorm"
)

type ICommentUsecase interface {
	CreateComment(comment model.Comment, visitedId uint) (model.CommentResponse, error)
}

type commentUsecase struct {
	cr repository.ICommentRepository
	cv validator.ICommentValidator
	nr repository.INotificationRepository
	db *gorm.DB
}

func NewCommentUsecase(
	cr repository.ICommentRepository,
	cv validator.ICommentValidator,
	nr repository.INotificationRepository,
	db *gorm.DB,
) ICommentUsecase {
	return &commentUsecase{cr, cv, nr, db}
}

func (cu *commentUsecase) CreateComment(comment model.Comment, visitedId uint) (model.CommentResponse, error) {
	if err := cu.cv.CommentValidate(comment); err != nil {
		return model.CommentResponse{}, err
	}

	tx := cu.db.Begin()
	if err := tx.Error; err != nil {
		return model.CommentResponse{}, err
	}

	if err := cu.cr.CreateComment(tx, &comment); err != nil {
		tx.Rollback()
		return model.CommentResponse{}, err
	}

	if comment.UserId != visitedId {
		notification := model.Notification{
			VisitorId: comment.UserId,
			VisitedId: visitedId,
			TweetId:   comment.TweetId,
			Action:    "comment",
			Read:      false,
		}

		if err := cu.nr.CreateNotification(tx, &notification); err != nil {
			tx.Rollback()
			return model.CommentResponse{}, err
		}
	}

	if err := tx.Commit().Error; err != nil {
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
