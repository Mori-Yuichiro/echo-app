package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	GetUserById(user *model.User, id uint) error
	CreateUser(user *model.User) error
	UpdateUser(user *model.User, userId uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserById(user *model.User, id uint) error {
	if err := ur.db.Preload("Tweets", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at desc")
	}).Preload("Tweets.User").Preload("Tweets.Favorites").Preload("Favorites", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at desc")
	}).Preload("Favorites.Tweet").Preload("Favorites.Tweet.User").Preload("Favorites.Tweet.Favorites").Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at desc")
	}).Preload("Comments.User").Where("id=?", id).Find(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) UpdateUser(user *model.User, userId uint) error {
	result := ur.db.Model(user).Clauses(clause.Returning{}).Where("id=?", userId).Updates(model.User{
		DisplayName:     user.DisplayName,
		Image:           user.Image,
		PhoneNumber:     user.PhoneNumber,
		Bio:             user.Bio,
		Location:        user.Location,
		Website:         user.Website,
		Birthday:        user.Birthday,
		ProfileImageUrl: user.ProfileImageUrl,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}
