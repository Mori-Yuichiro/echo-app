package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type INotificationRepository interface {
	GetNotificationsByUserId(notifications *[]model.Notification, userId uint) error
	CreateNotification(tx *gorm.DB, notification *model.Notification) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) INotificationRepository {
	return &notificationRepository{db}
}

func (nr *notificationRepository) GetNotificationsByUserId(notifications *[]model.Notification, userId uint) error {
	if err := nr.db.Preload("Visitor").Preload("Tweet").Where("visited_id=?", userId).Order("id DESC").Find(notifications).Error; err != nil {
		return err
	}
	return nil
}

func (nr *notificationRepository) CreateNotification(tx *gorm.DB, notification *model.Notification) error {
	if err := tx.Create(notification).Error; err != nil {
		return nil
	}
	return nil
}
