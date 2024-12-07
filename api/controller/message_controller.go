package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IMessageController interface {
	GetAllMessages(c echo.Context) error
	CreateMessage(c echo.Context) error
}

type messageController struct {
	mu usecase.IMessageUsecase
}

func NewMessageController(mu usecase.IMessageUsecase) IMessageController {
	return &messageController{mu}
}

func (mc *messageController) GetAllMessages(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	if userId == nil {
		return c.JSON(http.StatusInternalServerError, "you don't userId")
	}

	messageRes, err := mc.mu.GetAllMessages()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, messageRes)
}

func (mc *messageController) CreateMessage(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	floatId, err := strconv.ParseFloat(c.Param("roomId"), 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "param user_id can't be changed float64")
	}
	roomId := uint(floatId)

	message := model.Message{}
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	message.UserId = uint(userId.(float64))
	message.RoomId = roomId

	err = mc.mu.CreateMessage(message)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
