package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type IMessageUsecase interface {
	GetAllMessages() ([]model.MessageResponse, error)
	CreateMessage(message model.Message) error
}

type messageUsecase struct {
	mr repository.IMessageRepository
	mv validator.IMessageValidator
}

func NewMessageUsecase(mr repository.IMessageRepository, mv validator.IMessageValidator) IMessageUsecase {
	return &messageUsecase{mr, mv}
}

func (mu *messageUsecase) GetAllMessages() ([]model.MessageResponse, error) {
	messages := []model.Message{}
	if err := mu.mr.GetAllMessages(&messages); err != nil {
		return []model.MessageResponse{}, err
	}

	resMessages := []model.MessageResponse{}
	for _, message := range messages {
		user := model.UserResponse{
			ID:              message.User.ID,
			Email:           message.User.Email,
			Name:            message.User.Name,
			Image:           message.User.Image,
			DisplayName:     message.User.DisplayName,
			PhoneNumber:     message.User.PhoneNumber,
			Bio:             message.User.Bio,
			Location:        message.User.Location,
			Website:         message.User.Website,
			Birthday:        message.User.Birthday,
			ProfileImageUrl: message.User.ProfileImageUrl,
		}
		m := model.MessageResponse{
			ID:        message.ID,
			UserId:    message.UserId,
			RoomId:    message.RoomId,
			Message:   message.Message,
			CreatedAt: message.CreatedAt,
			UpdatedAt: message.UpdatedAt,
			User:      user,
		}

		resMessages = append(resMessages, m)
	}

	return resMessages, nil
}

func (mu *messageUsecase) CreateMessage(message model.Message) error {
	if err := mu.mv.MessageValidator(message); err != nil {
		return err
	}

	if err := mu.mr.CreateMessage(&message); err != nil {
		return err
	}

	return nil
}
