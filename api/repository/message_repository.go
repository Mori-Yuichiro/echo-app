package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IMessageRepository interface {
	GetAllMessages(messages *[]model.Message) error
	CreateMessage(message *model.Message) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) IMessageRepository {
	return &messageRepository{db}
}

func (mr *messageRepository) GetAllMessages(message *[]model.Message) error {
	if err := mr.db.Preload("User").Order("created_at DESC").Find(message).Error; err != nil {
		return err
	}
	return nil
}

func (mr *messageRepository) CreateMessage(message *model.Message) error {
	if err := mr.db.Create(message).Error; err != nil {
		return err
	}
	return nil
}
