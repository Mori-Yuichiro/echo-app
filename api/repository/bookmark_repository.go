package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IBookmarkRepository interface {
	GetAllBookmarks(bookmarks *[]model.Bookmark, userId uint) error
	CreateBookmark(bookmark *model.Bookmark) error
	DeleteBookmark(userId uint, tweetId uint) error
}

type bookmarkRepository struct {
	db *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) IBookmarkRepository {
	return &bookmarkRepository{db}
}

func (br *bookmarkRepository) GetAllBookmarks(bookmarks *[]model.Bookmark, userId uint) error {
	if err := br.db.Preload("Tweet").Preload("Tweet.User").Preload("Tweet.Favorites").Preload("Tweet.Retweets").Preload("Tweet.Bookmarks").Where("user_id=?", userId).Find(bookmarks).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookmarkRepository) CreateBookmark(bookmark *model.Bookmark) error {
	if err := br.db.Create(bookmark).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookmarkRepository) DeleteBookmark(userId uint, tweetId uint) error {
	result := br.db.Where("user_id=? AND tweet_id=?", userId, tweetId).Delete(&model.Bookmark{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
