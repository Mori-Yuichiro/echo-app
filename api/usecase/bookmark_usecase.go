package usecase

import (
	"encoding/json"
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IBookmarkUsecase interface {
	GetAllBookmarks(userId uint) ([]model.BookmarkResponse, error)
	CreateBookmark(bookmark model.Bookmark) (model.BookmarkResponse, error)
	DeleteBookmark(userId uint, tweetId uint) error
}

type bookmarkUsecase struct {
	br repository.IBookmarkRepository
}

func NewBookmarkUsecase(br repository.IBookmarkRepository) IBookmarkUsecase {
	return &bookmarkUsecase{br}
}

func (bu *bookmarkUsecase) GetAllBookmarks(userId uint) ([]model.BookmarkResponse, error) {
	bookmarks := []model.Bookmark{}
	if err := bu.br.GetAllBookmarks(&bookmarks, userId); err != nil {
		return []model.BookmarkResponse{}, err
	}

	resBookmarks := []model.BookmarkResponse{}
	for _, v := range bookmarks {
		// bookmarkしたtweetが持つfavoriteデータの取得
		var favorites []model.FavoriteResponse
		for _, fav := range v.Tweet.Favorites {
			favorites = append(favorites, model.FavoriteResponse{
				ID:        fav.ID,
				UserId:    fav.UserId,
				TweetId:   fav.TweetId,
				CreatedAt: fav.CreatedAt,
				UpdatedAt: fav.UpdatedAt,
			})
		}

		// bookmarkしたtweetが持つretweetデータの取得
		var retweets []model.RetweetResponse
		for _, ret := range v.Tweet.Retweets {
			retweets = append(retweets, model.RetweetResponse{
				ID:        ret.ID,
				UserId:    ret.UserId,
				TweetId:   ret.TweetId,
				CreatedAt: ret.CreatedAt,
				UpdatedAt: ret.UpdatedAt,
			})
		}

		var bookmarks []model.BookmarkResponse
		for _, book := range v.Tweet.Bookmarks {
			bookmarks = append(bookmarks, model.BookmarkResponse{
				ID:        book.ID,
				UserId:    book.UserId,
				TweetId:   book.TweetId,
				CreatedAt: book.CreatedAt,
				UpdatedAt: book.UpdatedAt,
			})
		}

		// bookmarkしたtweetのimage
		var image_urls []string
		if v.Tweet.ImageUrls != "" {
			err := json.Unmarshal([]byte(v.Tweet.ImageUrls), &image_urls)
			if err != nil {
				return []model.BookmarkResponse{}, err
			}
		}

		tweet := model.TweetResponse{
			ID:        v.Tweet.ID,
			Content:   v.Tweet.Content,
			ImageUrls: image_urls,
			CreatedAt: v.Tweet.CreatedAt,
			UpdatedAt: v.Tweet.UpdatedAt,
			User:      v.Tweet.User,
			Favorites: favorites,
			Retweets:  retweets,
			Bookmarks: bookmarks,
		}

		bookmark := model.BookmarkResponse{
			ID:        v.ID,
			UserId:    v.UserId,
			TweetId:   v.TweetId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Tweet:     tweet,
		}
		resBookmarks = append(resBookmarks, bookmark)
	}
	return resBookmarks, nil
}

func (bu *bookmarkUsecase) CreateBookmark(bookmark model.Bookmark) (model.BookmarkResponse, error) {
	if err := bu.br.CreateBookmark(&bookmark); err != nil {
		return model.BookmarkResponse{}, err
	}

	resBookmark := model.BookmarkResponse{
		ID:        bookmark.ID,
		UserId:    bookmark.UserId,
		TweetId:   bookmark.TweetId,
		CreatedAt: bookmark.CreatedAt,
		UpdatedAt: bookmark.UpdatedAt,
	}
	return resBookmark, nil
}

func (bu *bookmarkUsecase) DeleteBookmark(userId uint, tweetId uint) error {
	if err := bu.br.DeleteBookmark(userId, tweetId); err != nil {
		return err
	}
	return nil
}
