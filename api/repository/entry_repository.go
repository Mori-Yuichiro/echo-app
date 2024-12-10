package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IEntryRepository interface {
	GetEntryByUserId(entry *[]model.Entry, userId uint) error
	GetEntryByRoomAndUserId(entry *model.Entry, userId uint, roomId uint) error
	CreateEntry(entry *model.Entry) error
}

type entryRepository struct {
	db *gorm.DB
}

func NewEntryRepository(db *gorm.DB) IEntryRepository {
	return &entryRepository{db}
}

func (er *entryRepository) GetEntryByUserId(entries *[]model.Entry, userId uint) error {
	if err := er.db.Preload("User").Where("user_id=?", userId).Find(entries).Error; err != nil {
		return err
	}
	return nil
}

func (er *entryRepository) GetEntryByRoomAndUserId(entry *model.Entry, userId uint, roomId uint) error {
	if err := er.db.Preload("User").Where("room_id=?", roomId).Not("user_id=?", userId).First(entry).Error; err != nil {
		return err
	}
	return nil
}

func (er *entryRepository) CreateEntry(entry *model.Entry) error {
	if err := er.db.Create(entry).Error; err != nil {
		return err
	}
	return nil
}
