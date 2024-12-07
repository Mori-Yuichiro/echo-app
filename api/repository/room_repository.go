package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IRoomRepository interface {
	CreateRoom(room *model.Room) error
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) IRoomRepository {
	return &roomRepository{db}
}

func (rr *roomRepository) CreateRoom(room *model.Room) error {
	if err := rr.db.Create(room).Error; err != nil {
		return err
	}
	return nil
}
