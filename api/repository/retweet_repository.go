package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IRetweetRepository interface {
	CreateRetweet(retweet *model.Retweet) error
	DeleteRetweet(userId uint, tweetId uint) error
}

type retweetRepository struct {
	db *gorm.DB
}

func NewRetweetRepository(db *gorm.DB) IRetweetRepository {
	return &retweetRepository{db}
}

func (rr *retweetRepository) CreateRetweet(retweet *model.Retweet) error {
	if err := rr.db.Create(retweet).Error; err != nil {
		return err
	}
	return nil
}

func (rr *retweetRepository) DeleteRetweet(userId uint, tweetId uint) error {
	result := rr.db.Where("user_id=? AND tweet_id=?", userId, tweetId).Delete(&model.Retweet{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
