package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IBookmarkUsecase interface {
	CreateBookmark(bookmark model.Bookmark) (model.BookmarkResponse, error)
	DeleteBookmark(userId uint, tweetId uint) error
}

type bookmarkUsecase struct {
	br repository.IBookmarkRepository
}

func NewBookmarkUsecase(br repository.IBookmarkRepository) IBookmarkUsecase {
	return &bookmarkUsecase{br}
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
