package controller

import (
	"go-rest-api/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type INotificationController interface {
	GetNotificationsByUserId(c echo.Context) error
}

type notificationController struct {
	nu usecase.INotificationUsecase
}

func NewNotificationController(nu usecase.INotificationUsecase) INotificationController {
	return &notificationController{nu}
}

func (nc *notificationController) GetNotificationsByUserId(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	notificationRes, err := nc.nu.GetNotificationsByUserId(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, notificationRes)
}
