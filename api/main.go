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
	commentValidator := validator.NewCommentValidator()
	messageValidator := validator.NewMessageValidator()

	notificationRepository := repository.NewNotificationRepository(db)
	notificationUsecase := usecase.NewNotificationUsecase(notificationRepository)
	notificationController := controller.NewNotificationController(notificationUsecase)

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	imageRepository := repository.NewImageRepository(cld)
	imageUsecase := usecase.NewImageUsecase(imageRepository)
	imageController := controller.NewImageController(imageUsecase)

	tweetRepository := repository.NewTweetRepository(db)
	tweetUsecase := usecase.NewTweetUsecase(tweetRepository, tweetValidator)
	tweetController := controller.NewTweetController(tweetUsecase)

	favoriteRepository := repository.NewFavoriteRepository(db)
	favoriteUsecase := usecase.NewFavoriteUsecase(favoriteRepository, notificationRepository, db)
	favoriteController := controller.NewFavoriteController(favoriteUsecase)

	commentRepository := repository.NewCommentRepository(db)
	commentUsecase := usecase.NewCommentUsecase(
		commentRepository,
		commentValidator,
		notificationRepository,
		db,
	)
	commentController := controller.NewCommentController(commentUsecase)

	retweetRepository := repository.NewRetweetRepository(db)
	retweetUsecase := usecase.NewRetweetUsecase(retweetRepository, notificationRepository, db)
	retweetController := controller.NewRetweetController(retweetUsecase)

	bookmarkRepository := repository.NewBookmarkRepository(db)
	bookmarkUsecase := usecase.NewBookmarkUsecase(bookmarkRepository)
	bookmarkController := controller.NewBookmarkController(bookmarkUsecase)

	relationshipRepository := repository.NewRelationshipRepository(db)
	relationshipUsecase := usecase.NewRelationshipUsecase(relationshipRepository)
	relationshipController := controller.NewRelationshipController(relationshipUsecase)

	roomRepository := repository.NewRoomRepository(db)
	roomUsecase := usecase.NewRoomUsecase(roomRepository)
	roomController := controller.NewRoomController(roomUsecase)

	entryRepository := repository.NewEntryRepository(db)
	entryUsecase := usecase.NewEntryUsecase(entryRepository)
	entryController := controller.NewEntryController(entryUsecase)

	messageRepository := repository.NewMessageRepository(db)
	messageUsecase := usecase.NewMessageUsecase(messageRepository, messageValidator)
	messageController := controller.NewMessageController(messageUsecase)

	e := router.NewRouter(
		userController,
		imageController,
		tweetController,
		favoriteController,
		commentController,
		retweetController,
		bookmarkController,
		relationshipController,
		roomController,
		entryController,
		messageController,
		notificationController,
	)
	e.Logger.Fatal(e.Start(":8080"))
}
