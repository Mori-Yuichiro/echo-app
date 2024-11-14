package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IFavoriteController interface {
	CreateFavorite(c echo.Context) error
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

	favorite := model.Favorite{}
	if err := c.Bind(&favorite); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	favorite.UserId = uint(userId.(float64))
	favoriteRes, err := fc.fu.CreateFavorite(favorite)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, favoriteRes)
}
