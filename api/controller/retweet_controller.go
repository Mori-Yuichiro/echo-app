package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IRetweetController interface {
	CreateRetweet(c echo.Context) error
	DeleteRetweet(c echo.Context) error
}

type retweetController struct {
	ru usecase.IRetweetUsecase
}

func NewRetweetController(ru usecase.IRetweetUsecase) IRetweetController {
	return &retweetController{ru}
}

func (rc *retweetController) CreateRetweet(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	if userId == nil {
		return c.JSON(http.StatusInternalServerError, "you don't have userId")
	}

	id := c.Param("tweetId")
	tweetId, _ := strconv.Atoi(id)
	retweet := model.Retweet{
		UserId:  uint(userId.(float64)),
		TweetId: uint(tweetId),
	}

	retweetRes, err := rc.ru.CreateRetweet(retweet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, retweetRes)
}

func (rc *retweetController) DeleteRetweet(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := c.Param("tweetId")
	tweetId, _ := strconv.Atoi(id)

	err := rc.ru.DeleteRetweet(uint(userId.(float64)), uint(tweetId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
