package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	// GetAllComments(comment *[]model.Comment) error
	CreateComment(comment *model.Comment) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &commentRepository{db}
}

// func (cr *commentRepository) GetAllComments(comment *[]model.Comment) error {}

func (cr *commentRepository) CreateComment(comment *model.Comment) error {
	if err := cr.db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}
