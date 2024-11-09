package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IImageController interface {
	UploadImage(c echo.Context) error
}

type imageController struct {
	iu usecase.IImageUsecase
}

func NewImageController(iu usecase.IImageUsecase) IImageController {
	return &imageController{iu}
}

func (ic *imageController) UploadImage(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	if userId == nil {
		return c.JSON(http.StatusInternalServerError, "you don't userId")
	}

	image := model.Image{}
	if err := c.Bind(&image); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	imageRes, err := ic.iu.UploadImage(image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, imageRes)
}
