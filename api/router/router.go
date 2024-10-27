package router

import (
	"go-rest-api/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, tc controller.ITweetController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	t := e.Group("/tweets")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTweets)
	t.GET("/:tweetId", tc.GetTweetById)
	t.POST("", tc.CreateTweet)
	t.DELETE("/:tweetId", tc.DeleteTweet)

	return e
}
