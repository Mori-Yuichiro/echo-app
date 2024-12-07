package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IMessageValidator interface {
	MessageValidator(message model.Message) error
}

type messageValidator struct{}

func NewMessageValidator() IMessageValidator {
	return &messageValidator{}
}

func (mv *messageValidator) MessageValidator(message model.Message) error {
	return validation.ValidateStruct(
		&message,
		validation.Field(
			&message.Message,
			validation.Required.Error("message is required"),
			validation.RuneLength(1, 140).Error("message limit min 1 max 140 chars"),
		),
	)
}
