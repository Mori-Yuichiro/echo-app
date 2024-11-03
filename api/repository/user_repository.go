package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUesrByEmail(user *model.User, email string) error
	GetUesrById(user *model.User, id uint) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUesrByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUesrById(user *model.User, id uint) error {
	if err := ur.db.Where("id=?", id).First(user).Error; err != nil {
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
