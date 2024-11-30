package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IRelationshipRepository interface {
	CreateRelationship(relationship *model.Relationship) error
	DeleteRelationship(followerId uint, followedId uint) error
}

type relationshipRepository struct {
	db *gorm.DB
}

func NewRelationshipRepository(db *gorm.DB) IRelationshipRepository {
	return &relationshipRepository{db}
}

func (rr *relationshipRepository) CreateRelationship(relationship *model.Relationship) error {
	if err := rr.db.Create(relationship).Error; err != nil {
		return err
	}
	return nil
}

func (rr *relationshipRepository) DeleteRelationship(followerId uint, followedId uint) error {
	result := rr.db.Where("follower_id=? AND followed_id=?", followerId, followedId).Delete(&model.Relationship{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object doed not exist")
	}
	return nil
}
