package usecase

import (
	"encoding/json"
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
	GetUserById(id uint) (model.UserResponse, error)
	UpdateUser(user model.User, userId uint) (model.UserResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(
	ur repository.IUserRepository,
	uv validator.IUserValidator,
) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.SignUpValidate(user); err != nil {
		return model.UserResponse{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Name: user.Name, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:              newUser.ID,
		Email:           newUser.Email,
		Name:            newUser.Name,
		Image:           newUser.Image,
		DisplayName:     newUser.DisplayName,
		PhoneNumber:     newUser.PhoneNumber,
		Bio:             newUser.Bio,
		Location:        newUser.Location,
		Website:         newUser.Website,
		Birthday:        newUser.Birthday,
		ProfileImageUrl: newUser.ProfileImageUrl,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.LogInValidate(user); err != nil {
		return "", err
	}

	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uu *userUsecase) GetUserById(id uint) (model.UserResponse, error) {
	user := model.User{}
	if err := uu.ur.GetUserById(&user, id); err != nil {
		return model.UserResponse{}, err
	}

	// Userがツイートした内容
	var userTweets []model.TweetResponse
	for _, v := range user.Tweets {
		var image_urls []string
		if v.ImageUrls != "" {
			err := json.Unmarshal([]byte(v.ImageUrls), &image_urls)
			if err != nil {
				return model.UserResponse{}, err
			}
		}

		// Userが投稿したTweetのFavoritesを取得
		var tweetFavorite []model.FavoriteResponse
		for _, fav := range v.Favorites {
			tweetFavorite = append(tweetFavorite, model.FavoriteResponse{
				ID:        fav.ID,
				UserId:    fav.UserId,
				TweetId:   fav.TweetId,
				CreatedAt: fav.CreatedAt,
				UpdatedAt: fav.UpdatedAt,
			})
		}

		// userが投稿したtweetのretweet数を取得
		var tweetRetweets []model.RetweetResponse
		for _, ret := range v.Retweets {
			tweetRetweets = append(tweetRetweets, model.RetweetResponse{
				ID:        ret.ID,
				UserId:    ret.UserId,
				TweetId:   ret.TweetId,
				CreatedAt: ret.CreatedAt,
				UpdatedAt: ret.UpdatedAt,
			})
		}

		userTweet := model.TweetResponse{
			ID:        v.ID,
			Content:   v.Content,
			ImageUrls: image_urls,
			User:      v.User,
			Favorites: tweetFavorite,
			Retweets:  tweetRetweets,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		userTweets = append(userTweets, userTweet)
	}

	// Userがいいねしたツイート
	var favorites []model.FavoriteResponse
	for _, fav := range user.Favorites {
		var image_urls []string
		if fav.Tweet.ImageUrls != "" {
			err := json.Unmarshal([]byte(fav.Tweet.ImageUrls), &image_urls)
			if err != nil {
				return model.UserResponse{}, err
			}
		}

		// userがfavoriteしたtweetが保持するfavorite数
		var favTweetFavorites []model.FavoriteResponse
		for _, favTweetFavorite := range fav.Tweet.Favorites {
			favTweetFavorites = append(favTweetFavorites, model.FavoriteResponse{
				ID:        favTweetFavorite.ID,
				UserId:    favTweetFavorite.UserId,
				TweetId:   favTweetFavorite.TweetId,
				CreatedAt: favTweetFavorite.CreatedAt,
				UpdatedAt: favTweetFavorite.UpdatedAt,
			})
		}

		// userがfavoriteしたtweetが保持するretweet数
		var favTweetRetweets []model.RetweetResponse
		for _, favTweetRetweet := range fav.Tweet.Retweets {
			favTweetRetweets = append(favTweetRetweets, model.RetweetResponse{
				ID:        favTweetRetweet.ID,
				UserId:    favTweetRetweet.UserId,
				TweetId:   favTweetRetweet.TweetId,
				CreatedAt: favTweetRetweet.CreatedAt,
				UpdatedAt: favTweetRetweet.UpdatedAt,
			})
		}

		favTweet := model.TweetResponse{
			ID:        fav.Tweet.ID,
			Content:   fav.Tweet.Content,
			ImageUrls: image_urls,
			User:      fav.Tweet.User,
			CreatedAt: fav.Tweet.CreatedAt,
			UpdatedAt: fav.Tweet.UpdatedAt,
			Favorites: favTweetFavorites,
			Retweets:  favTweetRetweets,
		}

		favorite := model.FavoriteResponse{
			ID:        fav.ID,
			UserId:    fav.UserId,
			TweetId:   fav.TweetId,
			CreatedAt: fav.CreatedAt,
			UpdatedAt: fav.UpdatedAt,
			Tweet:     favTweet,
		}
		favorites = append(favorites, favorite)
	}

	// Userが投稿したCommentsを取得
	var comments []model.CommentResponse
	for _, com := range user.Comments {
		comments = append(comments, model.CommentResponse{
			ID:        com.ID,
			Comment:   com.Comment,
			UserId:    com.UserId,
			TweetId:   com.TweetId,
			CreatedAt: com.CreatedAt,
			UpdatedAt: com.UpdatedAt,
			User:      com.User,
		})
	}

	// Userがリツイートした内容
	var retweets []model.RetweetResponse
	for _, ret := range user.Retweets {
		var image_urls []string
		if ret.Tweet.ImageUrls != "" {
			err := json.Unmarshal([]byte(ret.Tweet.ImageUrls), &image_urls)
			if err != nil {
				return model.UserResponse{}, err
			}
		}

		// userがretweetしたtweetが保持するfavorite数
		var retweetTweetFavorites []model.FavoriteResponse
		for _, retTweetFavorite := range ret.Tweet.Favorites {
			retweetTweetFavorites = append(retweetTweetFavorites, model.FavoriteResponse{
				ID:        retTweetFavorite.ID,
				UserId:    retTweetFavorite.UserId,
				TweetId:   retTweetFavorite.TweetId,
				CreatedAt: retTweetFavorite.CreatedAt,
				UpdatedAt: retTweetFavorite.UpdatedAt,
			})
		}

		// userがretweetしたtweetが保持するretweet数
		var retweetTweetRetweets []model.RetweetResponse
		for _, retTweetRetweet := range ret.Tweet.Retweets {
			retweetTweetRetweets = append(retweetTweetRetweets, model.RetweetResponse{
				ID:        retTweetRetweet.ID,
				UserId:    retTweetRetweet.UserId,
				TweetId:   retTweetRetweet.TweetId,
				CreatedAt: retTweetRetweet.CreatedAt,
				UpdatedAt: retTweetRetweet.UpdatedAt,
			})
		}

		retTweet := model.TweetResponse{
			ID:        ret.Tweet.ID,
			Content:   ret.Tweet.Content,
			ImageUrls: image_urls,
			User:      ret.Tweet.User,
			CreatedAt: ret.Tweet.CreatedAt,
			UpdatedAt: ret.Tweet.UpdatedAt,
			Favorites: retweetTweetFavorites,
			Retweets:  retweetTweetRetweets,
		}

		retweet := model.RetweetResponse{
			ID:        ret.ID,
			UserId:    ret.UserId,
			TweetId:   ret.TweetId,
			CreatedAt: ret.CreatedAt,
			UpdatedAt: ret.UpdatedAt,
			Tweet:     retTweet,
		}
		retweets = append(retweets, retweet)
	}

	// userのfollowerデータ
	var followers []model.RelationshipResponse
	for _, follower := range user.Followers {
		followers = append(followers, model.RelationshipResponse{
			ID:         follower.ID,
			FollowerId: follower.FollowerId,
			FollowedId: follower.FollowedId,
			CreatedAt:  follower.CreatedAt,
			UpdatedAt:  follower.UpdatedAt,
		})
	}

	// userのfollowedデータ
	var followeds []model.RelationshipResponse
	for _, followed := range user.Followeds {
		followeds = append(followeds, model.RelationshipResponse{
			ID:         followed.ID,
			FollowerId: followed.FollowerId,
			FollowedId: followed.FollowedId,
			CreatedAt:  followed.CreatedAt,
			UpdatedAt:  followed.UpdatedAt,
		})
	}

	resUser := model.UserResponse{
		ID:              user.ID,
		Email:           user.Email,
		Name:            user.Name,
		Image:           user.Image,
		DisplayName:     user.DisplayName,
		PhoneNumber:     user.PhoneNumber,
		Bio:             user.Bio,
		Location:        user.Location,
		Website:         user.Website,
		Birthday:        user.Birthday,
		ProfileImageUrl: user.ProfileImageUrl,
		Tweets:          userTweets,
		Favorites:       favorites,
		Comments:        comments,
		Retweets:        retweets,
		Followers:       followers,
		Followeds:       followeds,
	}
	return resUser, nil
}

func (uu *userUsecase) UpdateUser(user model.User, userId uint) (model.UserResponse, error) {
	if err := uu.ur.UpdateUser(&user, userId); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:              user.ID,
		Email:           user.Email,
		Name:            user.Name,
		Image:           user.Image,
		DisplayName:     user.DisplayName,
		PhoneNumber:     user.PhoneNumber,
		Bio:             user.Bio,
		Location:        user.Location,
		Website:         user.Website,
		Birthday:        user.Birthday,
		ProfileImageUrl: user.ProfileImageUrl,
	}

	return resUser, nil
}
