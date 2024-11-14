package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IFavoriteController interface {
	CreateFavorite(c echo.Context) error
	DeleteFavorite(c echo.Context) error
}

type favoriteController struct {
	fu usecase.IFavoriteUsecase
}

func NewFavoriteController(fu usecase.IFavoriteUsecase) IFavoriteController {
	return &favoriteController{fu}
}

func (fc *favoriteController) CreateFavorite(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	if userId == nil {
		return c.JSON(http.StatusInternalServerError, "you don't have userId")
	}

	id := c.Param("tweetId")
	tweetId, _ := strconv.Atoi(id)
	favorite := model.Favorite{
		UserId:  uint(userId.(float64)),
		TweetId: uint(tweetId),
	}

	favoriteRes, err := fc.fu.CreateFavorite(favorite)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, favoriteRes)
}

func (fc *favoriteController) DeleteFavorite(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := c.Param("tweetId")
	tweetId, _ := strconv.Atoi(id)

	err := fc.fu.DeleteFavorite(uint(userId.(float64)), uint(tweetId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
