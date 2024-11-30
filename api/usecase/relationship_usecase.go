package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IRelationshipUsecase interface {
	CreateRelationship(relationship model.Relationship) (model.RelationshipResponse, error)
	DeleteRelationship(followerId uint, followedId uint) error
}

type relationshipUsecase struct {
	rr repository.IRelationshipRepository
}

func NewRelationshipUsecase(rr repository.IRelationshipRepository) IRelationshipUsecase {
	return &relationshipUsecase{rr}
}

func (ru *relationshipUsecase) CreateRelationship(relationship model.Relationship) (model.RelationshipResponse, error) {
	if err := ru.rr.CreateRelationship(&relationship); err != nil {
		return model.RelationshipResponse{}, err
	}

	resRelationship := model.RelationshipResponse{
		ID:         relationship.ID,
		FollowerId: relationship.FollowerId,
		FollowedId: relationship.FollowedId,
		CreatedAt:  relationship.CreatedAt,
		UpdatedAt:  relationship.UpdatedAt,
	}
	return resRelationship, nil
}

func (ru *relationshipUsecase) DeleteRelationship(followerId uint, followedId uint) error {
	if err := ru.rr.DeleteRelationship(followerId, followedId); err != nil {
		return err
	}
	return nil
}
