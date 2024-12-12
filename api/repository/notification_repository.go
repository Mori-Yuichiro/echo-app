package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type INotificationRepository interface {
	CreateNotification(notification *model.Notification) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) INotificationRepository {
	return &notificationRepository{db}
}

func (nr *notificationRepository) CreateNotification(notification *model.Notification) error {
	if err := nr.db.Create(notification).Error; err != nil {
		return nil
	}
	return nil
}
