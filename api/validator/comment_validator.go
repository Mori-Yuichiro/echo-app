package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ICommentValidator interface {
	CommentValidate(comment model.Comment) error
}

type commentValidator struct{}

func NewCommentValidator() ICommentValidator {
	return &commentValidator{}
}

func (cv *commentValidator) CommentValidate(comment model.Comment) error {
	return validation.ValidateStruct(&comment,
		validation.Field(
			&comment.Comment,
			validation.Required.Error("comment is required"),
			validation.RuneLength(1, 140).Error("comment limit min 1 max 140 chars"),
		),
	)
}
