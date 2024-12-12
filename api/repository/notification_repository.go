package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type INotificationRepository interface {
	CreateNotification(tx *gorm.DB, notification *model.Notification) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) INotificationRepository {
	return &notificationRepository{db}
}

func (nr *notificationRepository) CreateNotification(tx *gorm.DB, notification *model.Notification) error {
	if err := tx.Create(notification).Error; err != nil {
		return nil
	}
	return nil
}
