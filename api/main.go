package main

import (
	"fmt"
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func main() {
	db := db.NewDB()
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
	// cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		fmt.Println(err.Error())
	}

	userValidator := validator.NewUserValidator()
	tweetValidator := validator.NewTweetValidator()

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	imageRepository := repository.NewImageRepository(cld)
	imageUsecase := usecase.NewImageUsecase(imageRepository)
	imageController := controller.NewImageController(imageUsecase)

	tweetRepository := repository.NewTweetRepository(db)
	tweetUsecase := usecase.NewTweetUsecase(tweetRepository, tweetValidator)
	tweetController := controller.NewTweetController(tweetUsecase)

	e := router.NewRouter(userController, imageController, tweetController)
	e.Logger.Fatal(e.Start(":8080"))
}
