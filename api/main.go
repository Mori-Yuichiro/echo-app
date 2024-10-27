package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

func main() {
	db := db.NewDB()

	userValidator := validator.NewUserValidator()
	tweetValidator := validator.NewTweetValidator()

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	tweetRepository := repository.NewTweetRepository(db)
	tweetUsecase := usecase.NewTweetUsecase(tweetRepository, tweetValidator)
	tweetController := controller.NewTweetController(tweetUsecase)

	e := router.NewRouter(userController, tweetController)
	e.Logger.Fatal(e.Start(":8080"))
}
