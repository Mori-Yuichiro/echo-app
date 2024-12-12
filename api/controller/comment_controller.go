package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ICommentController interface {
	CreateComment(c echo.Context) error
}

type commentController struct {
	cu usecase.ICommentUsecase
}

func NewCommentController(cu usecase.ICommentUsecase) ICommentController {
	return &commentController{cu}
}

func (cc *commentController) CreateComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	comment := model.Comment{}
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	visitedId, err := strconv.ParseFloat(c.Param("visitedId"), 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	comment.UserId = uint(userId.(float64))
	commentRes, err := cc.cu.CreateComment(comment, uint(visitedId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, commentRes)
}
