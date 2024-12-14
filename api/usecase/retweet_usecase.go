package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"

	"gorm.io/gorm"
)

type IRetweetUsecase interface {
	CreateRetweet(retweet model.Retweet, visitedId uint) (model.RetweetResponse, error)
	DeleteRetweet(userId uint, tweetId uint) error
}

type retweetUsecase struct {
	rr repository.IRetweetRepository
	nr repository.INotificationRepository
	db *gorm.DB
}

func NewRetweetUsecase(
	rr repository.IRetweetRepository,
	nr repository.INotificationRepository,
	db *gorm.DB,
) IRetweetUsecase {
	return &retweetUsecase{rr, nr, db}
}

func (ru *retweetUsecase) CreateRetweet(retweet model.Retweet, visitedId uint) (model.RetweetResponse, error) {
	tx := ru.db.Begin()
	if err := tx.Error; err != nil {
		return model.RetweetResponse{}, err
	}

	if err := ru.rr.CreateRetweet(tx, &retweet); err != nil {
		tx.Rollback()
		return model.RetweetResponse{}, err
	}

	if retweet.UserId != visitedId {
		notification := model.Notification{
			VisitorId: retweet.UserId,
			VisitedId: visitedId,
			TweetId:   nil,
			Action:    "retweet",
			Read:      false,
		}

		if err := ru.nr.CreateNotification(tx, &notification); err != nil {
			tx.Rollback()
			return model.RetweetResponse{}, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return model.RetweetResponse{}, err
	}

	resRetweet := model.RetweetResponse{
		ID:        retweet.ID,
		UserId:    retweet.UserId,
		TweetId:   retweet.TweetId,
		CreatedAt: retweet.CreatedAt,
		UpdatedAt: retweet.UpdatedAt,
	}
	return resRetweet, nil
}

func (ru *retweetUsecase) DeleteRetweet(userId uint, tweetId uint) error {
	if err := ru.rr.DeleteRetweet(userId, tweetId); err != nil {
		return err
	}
	return nil
}
