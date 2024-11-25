package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IBookmarkController interface {
	CreateBookmark(c echo.Context) error
	DeleteBookmark(c echo.Context) error
}

type bookmarkController struct {
	bu usecase.IBookmarkUsecase
}

func NewBookmarkController(bu usecase.IBookmarkUsecase) IBookmarkController {
	return &bookmarkController{bu}
}

func (bc *bookmarkController) CreateBookmark(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := c.Param("tweetId")
	tweetId, _ := strconv.Atoi(id)
	bookmark := model.Bookmark{
		UserId:  uint(userId.(float64)),
		TweetId: uint(tweetId),
	}

	bookmarkRes, err := bc.bu.CreateBookmark(bookmark)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, bookmarkRes)
}

func (bc *bookmarkController) DeleteBookmark(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := c.Param("tweetId")
	tweetId, _ := strconv.Atoi(id)

	err := bc.bu.DeleteBookmark(uint(userId.(float64)), uint(tweetId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
