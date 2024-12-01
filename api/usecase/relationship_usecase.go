package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IRelationshipUsecase interface {
	GetFollowersById(userId uint) ([]model.RelationshipResponse, error)
	GetFollowedsById(userId uint) ([]model.RelationshipResponse, error)
	CreateRelationship(relationship model.Relationship) (model.RelationshipResponse, error)
	DeleteRelationship(followerId uint, followedId uint) error
}

type relationshipUsecase struct {
	rr repository.IRelationshipRepository
}

func NewRelationshipUsecase(rr repository.IRelationshipRepository) IRelationshipUsecase {
	return &relationshipUsecase{rr}
}

func (ru *relationshipUsecase) GetFollowersById(userId uint) ([]model.RelationshipResponse, error) {
	relationships := []model.Relationship{}
	if err := ru.rr.GetFollowersById(&relationships, userId); err != nil {
		return []model.RelationshipResponse{}, err
	}

	relResponse := []model.RelationshipResponse{}
	for _, rel := range relationships {
		follower := model.UserResponse{
			ID:              rel.Follower.ID,
			Email:           rel.Follower.Email,
			Name:            rel.Follower.Name,
			Image:           rel.Follower.Image,
			DisplayName:     rel.Follower.DisplayName,
			PhoneNumber:     rel.Follower.PhoneNumber,
			Bio:             rel.Follower.Bio,
			Location:        rel.Follower.Location,
			Website:         rel.Follower.Website,
			Birthday:        rel.Follower.Birthday,
			ProfileImageUrl: rel.Follower.ProfileImageUrl,
		}

		relResponse = append(relResponse, model.RelationshipResponse{
			ID:         rel.ID,
			FollowerId: rel.FollowerId,
			FollowedId: rel.FollowedId,
			Follower:   follower,
			CreatedAt:  rel.CreatedAt,
			UpdatedAt:  rel.UpdatedAt,
		})
	}

	return relResponse, nil
}

func (ru *relationshipUsecase) GetFollowedsById(userId uint) ([]model.RelationshipResponse, error) {
	relationships := []model.Relationship{}
	if err := ru.rr.GetFollowedsById(&relationships, userId); err != nil {
		return []model.RelationshipResponse{}, err
	}

	relResponse := []model.RelationshipResponse{}
	for _, rel := range relationships {
		followed := model.UserResponse{
			ID:              rel.Followed.ID,
			Email:           rel.Followed.Email,
			Name:            rel.Followed.Name,
			Image:           rel.Followed.Image,
			DisplayName:     rel.Followed.DisplayName,
			PhoneNumber:     rel.Followed.PhoneNumber,
			Bio:             rel.Followed.Bio,
			Location:        rel.Followed.Location,
			Website:         rel.Followed.Website,
			Birthday:        rel.Followed.Birthday,
			ProfileImageUrl: rel.Followed.ProfileImageUrl,
		}

		relResponse = append(relResponse, model.RelationshipResponse{
			ID:         rel.ID,
			FollowerId: rel.FollowerId,
			FollowedId: rel.FollowedId,
			Followed:   followed,
			CreatedAt:  rel.CreatedAt,
			UpdatedAt:  rel.UpdatedAt,
		})
	}

	return relResponse, nil
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
