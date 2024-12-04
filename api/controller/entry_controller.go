package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IEntryController interface {
	CreateEntry(c echo.Context) error
}

type entryController struct {
	eu usecase.IEntryUsecase
}

func NewEntryController(eu usecase.IEntryUsecase) IEntryController {
	return &entryController{eu}
}

func (ec *entryController) CreateEntry(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	entry := model.Entry{}
	if err := c.Bind(&entry); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	entry.UserId = uint(userId.(float64))
	err := ec.eu.CreateEntry(entry)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
