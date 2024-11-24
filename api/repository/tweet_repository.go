package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
)

type ITweetRepository interface {
	GetAllTweets(tweet *[]model.Tweet) error
	GetTweetById(tweet *model.Tweet, tweetId uint) error
	CreateTweet(tweet *model.Tweet) error
	DeleteTweet(userId uint, tweetId uint) error
}

type tweetRepository struct {
	db *gorm.DB
}

func NewTweetRepository(db *gorm.DB) ITweetRepository {
	return &tweetRepository{db}
}

func (tr *tweetRepository) GetAllTweets(tweet *[]model.Tweet) error {
	if err := tr.db.Preload("User").Preload("Favorites").Preload("Retweets").Order("created_at DESC").Find(tweet).Error; err != nil {
		return err
	}
	return nil
}

func (tr *tweetRepository) GetTweetById(tweet *model.Tweet, tweetId uint) error {
	if err := tr.db.Preload("User").Preload("Favorites").Preload("Comments").Preload("Comments.User").Order("created_at").First(tweet, tweetId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *tweetRepository) CreateTweet(tweet *model.Tweet) error {
	if err := tr.db.Create(tweet).Error; err != nil {
		return err
	}
	return nil
}

func (tr *tweetRepository) DeleteTweet(userId uint, tweetId uint) error {
	result := tr.db.Where("id=? AND user_id=?", tweetId, userId).Delete(&model.Tweet{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
