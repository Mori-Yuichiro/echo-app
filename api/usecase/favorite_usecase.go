package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"

	"gorm.io/gorm"
)

type IFavoriteUsecase interface {
	CreateFavorite(favorite model.Favorite, visitorId uint) (model.FavoriteResponse, error)
	DeleteFavorite(userId uint, tweetId uint) error
}

type favoriteUsecase struct {
	fr repository.IFavoriteRepository
	nr repository.INotificationRepository
	db *gorm.DB
}

func NewFavoriteUsecase(
	fr repository.IFavoriteRepository,
	nr repository.INotificationRepository,
	db *gorm.DB,
) IFavoriteUsecase {
	return &favoriteUsecase{fr, nr, db}
}

func (fu *favoriteUsecase) CreateFavorite(favorite model.Favorite, visitedId uint) (model.FavoriteResponse, error) {
	tx := fu.db.Begin()
	if err := tx.Error; err != nil {
		return model.FavoriteResponse{}, err
	}

	if err := fu.fr.CreateFavorite(tx, &favorite); err != nil {
		tx.Rollback()
		return model.FavoriteResponse{}, err
	}

	if favorite.UserId != visitedId {
		notification := model.Notification{
			VisitorId: favorite.UserId,
			VisitedId: visitedId,
			TweetId:   nil,
			Action:    "favorite",
			Read:      false,
		}

		if err := fu.nr.CreateNotification(tx, &notification); err != nil {
			tx.Rollback()
			return model.FavoriteResponse{}, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return model.FavoriteResponse{}, err
	}

	resFavorite := model.FavoriteResponse{
		ID:        favorite.ID,
		UserId:    favorite.UserId,
		TweetId:   favorite.TweetId,
		CreatedAt: favorite.CreatedAt,
		UpdatedAt: favorite.UpdatedAt,
	}
	return resFavorite, nil
}

func (fu *favoriteUsecase) DeleteFavorite(userId uint, tweetId uint) error {
	if err := fu.fr.DeleteFavorite(userId, tweetId); err != nil {
		return err
	}
	return nil
}
