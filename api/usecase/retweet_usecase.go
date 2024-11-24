package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IRetweetUsecase interface {
	CreateRetweet(retweet model.Retweet) (model.RetweetResponse, error)
	DeleteRetweet(userId uint, tweetId uint) error
}

type retweetUsecase struct {
	rr repository.IRetweetRepository
}

func NewRetweetUsecase(rr repository.IRetweetRepository) IRetweetUsecase {
	return &retweetUsecase{rr}
}

func (ru *retweetUsecase) CreateRetweet(retweet model.Retweet) (model.RetweetResponse, error) {
	if err := ru.rr.CreateRetweet(&retweet); err != nil {
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
