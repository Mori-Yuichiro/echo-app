package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type INotificationUsecase interface {
	GetNotificationsByUserId(userId uint) ([]model.NotificationResponse, error)
}

type notificaitonUsecase struct {
	nr repository.INotificationRepository
}

func NewNotificationUsecase(nr repository.INotificationRepository) INotificationUsecase {
	return &notificaitonUsecase{nr}
}

func (nu *notificaitonUsecase) GetNotificationsByUserId(userId uint) ([]model.NotificationResponse, error) {
	notifications := []model.Notification{}
	if err := nu.nr.GetNotificationsByUserId(&notifications, userId); err != nil {
		return []model.NotificationResponse{}, err
	}

	resNotifications := []model.NotificationResponse{}
	var user model.UserResponse
	for _, v := range notifications {
		user = model.UserResponse{
			ID:              v.Visitor.ID,
			Email:           v.Visitor.Email,
			Name:            v.Visitor.Name,
			Image:           v.Visitor.Image,
			DisplayName:     v.Visitor.DisplayName,
			PhoneNumber:     v.Visitor.PhoneNumber,
			Bio:             v.Visitor.Bio,
			Location:        v.Visitor.Location,
			Website:         v.Visitor.Website,
			Birthday:        v.Visitor.Birthday,
			ProfileImageUrl: v.Visitor.ProfileImageUrl,
		}

		notification := model.NotificationResponse{
			ID:        v.ID,
			VisitorId: v.VisitorId,
			VisitedId: v.VisitedId,
			TweetId:   v.TweetId,
			Action:    v.Action,
			Read:      v.Read,
			Visitor:   user,
		}
		resNotifications = append(resNotifications, notification)
	}
	return resNotifications, nil
}
