package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IFavoriteRepository interface {
	CreateFavorite(favorite *model.Favorite) error
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
