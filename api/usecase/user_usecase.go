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
		userTweet := model.TweetResponse{
			ID:        v.ID,
			Content:   v.Content,
			ImageUrls: image_urls,
			User:      v.User,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		userTweets = append(userTweets, userTweet)
	}

	// Userがいいねしたツイート
	var favoriteResponse []model.FavoriteResponse
	for _, fav := range user.Favorites {
		var image_urls []string
		if fav.Tweet.ImageUrls != "" {
			err := json.Unmarshal([]byte(fav.Tweet.ImageUrls), &image_urls)
			if err != nil {
				return model.UserResponse{}, err
			}
		}

		var favTweetFavorites []model.FavoriteResponse
		for _, tweetFav := range fav.Tweet.Favorites {
			favTweetFavorites = append(favTweetFavorites, model.FavoriteResponse{
				ID:        tweetFav.ID,
				UserId:    tweetFav.UserId,
				TweetId:   tweetFav.TweetId,
				CreatedAt: tweetFav.CreatedAt,
				UpdatedAt: tweetFav.UpdatedAt,
			})
		}

		favTweet := model.TweetResponse{
			ID:        fav.Tweet.ID,
			Content:   fav.Tweet.Content,
			ImageUrls: image_urls,
			User:      fav.Tweet.User,
			Favorites: favTweetFavorites,
			CreatedAt: fav.Tweet.CreatedAt,
			UpdatedAt: fav.Tweet.UpdatedAt,
		}

		favorite := model.FavoriteResponse{
			ID:        fav.ID,
			UserId:    fav.UserId,
			TweetId:   fav.TweetId,
			CreatedAt: fav.CreatedAt,
			UpdatedAt: fav.UpdatedAt,
			Tweet:     favTweet,
		}

		favoriteResponse = append(favoriteResponse, favorite)
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
		Favorites:       favoriteResponse,
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
