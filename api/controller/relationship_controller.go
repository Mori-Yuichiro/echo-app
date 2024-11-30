package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IRelationshipController interface {
	CreateRelationship(c echo.Context) error
	DeleteRelationship(c echo.Context) error
}

type relationshipController struct {
	ru usecase.IRelationshipUsecase
}

func NewRelationshipController(ru usecase.IRelationshipUsecase) IRelationshipController {
	return &relationshipController{ru}
}

func (rc *relationshipController) CreateRelationship(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := c.Param("userId")
	intId, _ := strconv.Atoi(id)
	followedId := uint(intId)
	followerId := uint(userId.(float64))

	relationship := model.Relationship{
		FollowerId: followerId,
		FollowedId: followedId,
	}

	relationshipRes, err := rc.ru.CreateRelationship(relationship)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, relationshipRes)
}

func (rc *relationshipController) DeleteRelationship(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := c.Param("userId")
	intId, _ := strconv.Atoi(id)
	followedId := uint(intId)
	followerId := uint(userId.(float64))

	err := rc.ru.DeleteRelationship(followerId, followedId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
