package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IRelationshipRepository interface {
	GetFollowersById(relationship *[]model.Relationship, userId uint) error
	GetFollowedsById(relationship *[]model.Relationship, userId uint) error
	CreateRelationship(relationship *model.Relationship) error
	DeleteRelationship(followerId uint, followedId uint) error
}

type relationshipRepository struct {
	db *gorm.DB
}

func NewRelationshipRepository(db *gorm.DB) IRelationshipRepository {
	return &relationshipRepository{db}
}

func (rr *relationshipRepository) GetFollowersById(relationship *[]model.Relationship, userId uint) error {
	if err := rr.db.Preload("Follower").Where("followed_id", userId).Find(relationship).Error; err != nil {
		return err
	}
	return nil
}

func (rr *relationshipRepository) GetFollowedsById(relationship *[]model.Relationship, userId uint) error {
	if err := rr.db.Preload("Followed").Where("follower_id", userId).Find(relationship).Error; err != nil {
		return err
	}
	return nil
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
