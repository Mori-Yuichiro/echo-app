package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IRoomUsecase interface {
	CreateRoom(room model.Room) (model.RoomResponse, error)
}

type roomUsecase struct {
	rr repository.IRoomRepository
}

func NewRoomUsecase(rr repository.IRoomRepository) IRoomUsecase {
	return &roomUsecase{rr}
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
