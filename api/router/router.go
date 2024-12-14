package router

import (
	"go-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	uc controller.IUserController,
	ic controller.IImageController,
	tc controller.ITweetController,
	fc controller.IFavoriteController,
	cc controller.ICommentController,
	rc controller.IRetweetController,
	bc controller.IBookmarkController,
	relc controller.IRelationshipController,
	rmc controller.IRoomController,
	ec controller.IEntryController,
	mc controller.IMessageController,
) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
			os.Getenv("FE_URL"),
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken,
		},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,

		// Postmanで動作確認をしたいからdefault
		// CookieSameSite: http.SameSiteDefaultMode,

		// CookieMaxAge: 60,
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	u := e.Group("/users")
	u.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	u.GET("", uc.GetUserIdByToken)
	u.GET("/:userId", uc.GetUserById)
	u.PUT("", uc.UpdateUser)

	// Follow
	r := u.Group("/:userId")
	r.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	r.POST("/follow", relc.CreateRelationship)
	r.DELETE("/follow", relc.DeleteRelationship)

	r.GET("/followers", relc.GetFollowersById)
	r.GET("/followeds", relc.GetFollowedsById)

	i := e.Group("/image-upload")
	i.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	i.POST("", ic.UploadImage)

	t := e.Group("/tweets")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTweets)
	t.GET("/:tweetId", tc.GetTweetById)
	t.POST("", tc.CreateTweet)
	t.DELETE("/:tweetId", tc.DeleteTweet)

	tid := t.Group("/:tweetId")
	tid.POST("/:visitedId/favorite", fc.CreateFavorite)
	tid.DELETE("/favorite", fc.DeleteFavorite)

	tid.POST("/:visitedId/retweet", rc.CreateRetweet)
	tid.DELETE("/retweet", rc.DeleteRetweet)

	tid.POST("/bookmark", bc.CreateBookmark)
	tid.DELETE("/bookmark", bc.DeleteBookmark)

	b := e.Group("/bookmarks")
	b.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	b.GET("", bc.GetAllBookmarks)

	c := e.Group("/:visitedId/comment")
	c.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	c.POST("", cc.CreateComment)

	// room
	room := e.Group("/room")
	room.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	room.GET("", rmc.GetRooms)
	room.POST("", rmc.CreateRoom)

	// entry
	entry := e.Group("/entry")
	entry.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	entry.GET("/:userId", ec.GetEntryByUserId)
	entry.GET("/room/:roomId", ec.GetEntryByRoomAndUserId)
	entry.POST("", ec.CreateEntry)

	// message
	m := room.Group("/:roomId")
	m.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	m.GET("/messages", mc.GetAllMessages)
	m.POST("/message", mc.CreateMessage)

	return e
}
