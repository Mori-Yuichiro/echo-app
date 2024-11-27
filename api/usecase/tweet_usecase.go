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
		//Tweetが持つFavoriteデータを取得
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

		// Tweetが持つRetweetを取得
		var retweets []model.RetweetResponse
		for _, ret := range v.Retweets {
			retweets = append(retweets, model.RetweetResponse{
				ID:        ret.ID,
				UserId:    ret.UserId,
				TweetId:   ret.TweetId,
				CreatedAt: ret.CreatedAt,
				UpdatedAt: ret.UpdatedAt,
			})
		}

		// tweetが持つbookmarkデータを取得
		var bookmarks []model.BookmarkResponse
		for _, book := range v.Bookmarks {
			bookmarks = append(bookmarks, model.BookmarkResponse{
				ID:        book.ID,
				UserId:    book.UserId,
				TweetId:   book.TweetId,
				CreatedAt: book.CreatedAt,
				UpdatedAt: book.UpdatedAt,
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
				Retweets:  retweets,
				Bookmarks: bookmarks,
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
				Retweets:  retweets,
				Bookmarks: bookmarks,
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

	var comments []model.CommentResponse
	for _, com := range tweet.Comments {
		comments = append(comments, model.CommentResponse{
			ID:        com.ID,
			Comment:   com.Comment,
			UserId:    com.UserId,
			TweetId:   com.TweetId,
			CreatedAt: com.CreatedAt,
			UpdatedAt: com.UpdatedAt,
			User:      com.User,
		})
	}

	var bookmarks []model.BookmarkResponse
	for _, book := range tweet.Bookmarks {
		bookmarks = append(bookmarks, model.BookmarkResponse{
			ID:        book.ID,
			UserId:    book.UserId,
			TweetId:   book.TweetId,
			CreatedAt: book.CreatedAt,
			UpdatedAt: book.UpdatedAt,
		})
	}

	resTweet := model.TweetResponse{
		ID:        tweet.ID,
		Content:   tweet.Content,
		ImageUrls: image_urls,
		User:      tweet.User,
		Favorites: favorites,
		Comments:  comments,
		Bookmarks: bookmarks,
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
