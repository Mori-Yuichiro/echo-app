package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IFavoriteUsecase interface {
	CreateFavorite(favorite model.Favorite) (model.FavoriteResponse, error)
}

type favoriteUsecase struct {
	fr repository.IFavoriteRepository
}

func NewFavoriteUsecase(fr repository.IFavoriteRepository) IFavoriteUsecase {
	return &favoriteUsecase{fr}
}

func (fu *favoriteUsecase) CreateFavorite(favorite model.Favorite) (model.FavoriteResponse, error) {
	if err := fu.fr.CreateFavorite(&favorite); err != nil {
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
