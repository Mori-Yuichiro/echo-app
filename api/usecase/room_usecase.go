package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IRoomUsecase interface {
	GetRooms() ([]model.RoomResponse, error)
	CreateRoom(room model.Room) (model.RoomResponse, error)
}

type roomUsecase struct {
	rr repository.IRoomRepository
}

func NewRoomUsecase(rr repository.IRoomRepository) IRoomUsecase {
	return &roomUsecase{rr}
}

func (ru *roomUsecase) GetRooms() ([]model.RoomResponse, error) {
	rooms := []model.Room{}
	if err := ru.rr.GetRooms(&rooms); err != nil {
		return []model.RoomResponse{}, err
	}

	resRooms := []model.RoomResponse{}
	for _, v := range rooms {
		entries := []model.EntryResponse{}
		user := model.UserResponse{}
		for _, entry := range v.Entries {
			user = model.UserResponse{
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

			entries = append(entries, model.EntryResponse{
				ID:        entry.ID,
				UserId:    entry.UserId,
				RoomId:    entry.RoomId,
				CreatedAt: entry.CreatedAt,
				UpdatedAt: entry.UpdatedAt,
				User:      user,
			})
		}

		messages := []model.MessageResponse{}
		for _, message := range v.Messages {
			messages = append(messages, model.MessageResponse{
				ID:        message.ID,
				UserId:    message.UserId,
				RoomId:    message.RoomId,
				Message:   message.Message,
				CreatedAt: message.CreatedAt,
				UpdatedAt: message.UpdatedAt,
			})
		}

		resRooms = append(resRooms, model.RoomResponse{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Entries:   entries,
			Messages:  messages,
		})
	}

	return resRooms, nil
}

func (ru *roomUsecase) CreateRoom(room model.Room) (model.RoomResponse, error) {
	if err := ru.rr.CreateRoom(&room); err != nil {
		return model.RoomResponse{}, err
	}
	resRoom := model.RoomResponse{
		ID:        room.ID,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}
	return resRoom, nil
}
