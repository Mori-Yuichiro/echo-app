package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IEntryUsecase interface {
	GetEntryByUserId(userId uint) ([]model.EntryResponse, error)
	GetEntryByRoomAndUserId(userId uint, roomId uint) (model.EntryResponse, error)
	CreateEntry(entry model.Entry) error
}

type entryUsecase struct {
	er repository.IEntryRepository
}

func NewEntryUsecase(er repository.IEntryRepository) IEntryUsecase {
	return &entryUsecase{er}
}

func (eu *entryUsecase) GetEntryByUserId(userId uint) ([]model.EntryResponse, error) {
	entries := []model.Entry{}
	if err := eu.er.GetEntryByUserId(&entries, userId); err != nil {
		return []model.EntryResponse{}, err
	}

	resEntry := []model.EntryResponse{}
	for _, entry := range entries {
		user := model.UserResponse{
			ID:              entry.User.ID,
			Email:           entry.User.Email,
			Name:            entry.User.Name,
			Image:           entry.User.Image,
			DisplayName:     entry.User.DisplayName,
			PhoneNumber:     entry.User.PhoneNumber,
			Bio:             entry.User.Bio,
			Location:        entry.User.Location,
			Website:         entry.User.Website,
			Birthday:        entry.User.Birthday,
			ProfileImageUrl: entry.User.ProfileImageUrl,
		}

		resEntry = append(resEntry, model.EntryResponse{
			ID:        entry.ID,
			UserId:    entry.UserId,
			RoomId:    entry.RoomId,
			CreatedAt: entry.CreatedAt,
			UpdatedAt: entry.UpdatedAt,
			User:      user,
		})
	}
	return resEntry, nil
}

func (eu *entryUsecase) GetEntryByRoomAndUserId(userId uint, roomId uint) (model.EntryResponse, error) {
	entry := model.Entry{}
	if err := eu.er.GetEntryByRoomAndUserId(&entry, userId, roomId); err != nil {
		return model.EntryResponse{}, err
	}

	user := model.UserResponse{
		ID:              entry.User.ID,
		Email:           entry.User.Email,
		Name:            entry.User.Name,
		Image:           entry.User.Image,
		DisplayName:     entry.User.DisplayName,
		PhoneNumber:     entry.User.PhoneNumber,
		Bio:             entry.User.Bio,
		Location:        entry.User.Location,
		Website:         entry.User.Website,
		Birthday:        entry.User.Birthday,
		ProfileImageUrl: entry.User.ProfileImageUrl,
	}

	resEntry := model.EntryResponse{
		ID:        entry.ID,
		UserId:    entry.UserId,
		RoomId:    entry.RoomId,
		CreatedAt: entry.CreatedAt,
		UpdatedAt: entry.UpdatedAt,
		User:      user,
	}
	return resEntry, nil
}

func (eu *entryUsecase) CreateEntry(entry model.Entry) error {
	if err := eu.er.CreateEntry(&entry); err != nil {
		return err
	}
	return nil
}
