package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(tx *gorm.DB, comment *model.Comment) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &commentRepository{db}
}

func (cr *commentRepository) CreateComment(tx *gorm.DB, comment *model.Comment) error {
	if err := tx.Create(comment).Error; err != nil {
		return err
	}
	return nil
}
