package usecase

import (
	"encoding/json"
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type ITweetUsecase interface {
	GetAllTweets() ([]model.TweetResponse, error)
	GetTweetById(tweetId uint) (model.TweetResponse, error)
	CreateTweet(tweet model.Tweet) (model.TweetResponse, error)
	DeleteTweet(userId uint, tweetId uint) error
}

type tweetUsecase struct {
	tr repository.ITweetRepository
	tv validator.ITweetValidator
}

func NewTweetUsecase(tr repository.ITweetRepository, tv validator.ITweetValidator) ITweetUsecase {
	return &tweetUsecase{tr, tv}
}

func (tu *tweetUsecase) GetAllTweets() ([]model.TweetResponse, error) {
	tweets := []model.Tweet{}
	if err := tu.tr.GetAllTweets(&tweets); err != nil {
		return []model.TweetResponse{}, err
	}

	resTweets := []model.TweetResponse{}
	for _, v := range tweets {
		var favorites []model.FavoriteResponse
		for _, fav := range v.Favorites {
			favorites = append(favorites, model.FavoriteResponse{
				ID:        fav.ID,
				UserId:    fav.UserId,
				TweetId:   fav.TweetId,
				CreatedAt: fav.CreatedAt,
				UpdatedAt: fav.UpdatedAt,
			})
		}

		if v.ImageUrls != "" {
			var image_urls []string
			err := json.Unmarshal([]byte(v.ImageUrls), &image_urls)
			if err != nil {
				return []model.TweetResponse{}, err
			}

			t := model.TweetResponse{
				ID:        v.ID,
				Content:   v.Content,
				ImageUrls: image_urls,
				User:      v.User,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				Favorites: favorites,
			}
			resTweets = append(resTweets, t)
		} else {
			t := model.TweetResponse{
				ID:        v.ID,
				Content:   v.Content,
				User:      v.User,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				Favorites: favorites,
			}
			resTweets = append(resTweets, t)
		}
	}
	return resTweets, nil
}

func (tu *tweetUsecase) GetTweetById(tweetId uint) (model.TweetResponse, error) {
	tweet := model.Tweet{}
	if err := tu.tr.GetTweetById(&tweet, tweetId); err != nil {
		return model.TweetResponse{}, err
	}

	var image_urls []string
	if tweet.ImageUrls != "" {
		err := json.Unmarshal([]byte(tweet.ImageUrls), &image_urls)
		if err != nil {
			return model.TweetResponse{}, err
		}
	}

	var favorites []model.FavoriteResponse
	for _, fav := range tweet.Favorites {
		favorites = append(favorites, model.FavoriteResponse{
			ID:        fav.ID,
			UserId:    fav.UserId,
			TweetId:   fav.TweetId,
			CreatedAt: fav.CreatedAt,
			UpdatedAt: fav.UpdatedAt,
		})
	}

	resTweet := model.TweetResponse{
		ID:        tweet.ID,
		Content:   tweet.Content,
		ImageUrls: image_urls,
		User:      tweet.User,
		Favorites: favorites,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
	}

	return resTweet, nil
}

func (tu *tweetUsecase) CreateTweet(tweet model.Tweet) (model.TweetResponse, error) {
	if err := tu.tv.TweetValidate(tweet); err != nil {
		return model.TweetResponse{}, err
	}

	if err := tu.tr.CreateTweet(&tweet); err != nil {
		return model.TweetResponse{}, err
	}

	resTweet := model.TweetResponse{
		ID:        tweet.ID,
		Content:   tweet.Content,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
	}
	return resTweet, nil
}

func (tu *tweetUsecase) DeleteTweet(userId uint, tweetId uint) error {
	if err := tu.tr.DeleteTweet(userId, tweetId); err != nil {
		return err
	}
	return nil
}
