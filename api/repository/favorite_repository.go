package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IFavoriteRepository interface {
	CreateFavorite(favorite *model.Favorite) error
	DeleteFavorite(userId uint, tweetId uint) error
}

type favoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) IFavoriteRepository {
	return &favoriteRepository{db}
}

func (fr *favoriteRepository) CreateFavorite(favorite *model.Favorite) error {
	if err := fr.db.Create(favorite).Error; err != nil {
		return err
	}
	return nil
}

func (fr *favoriteRepository) DeleteFavorite(userId uint, tweetId uint) error {
	result := fr.db.Where("user_id=? AND tweet_id=?", userId, tweetId).Delete(&model.Favorite{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
