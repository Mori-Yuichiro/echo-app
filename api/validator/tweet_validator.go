package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITweetValidator interface {
	TweetValidate(tweet model.Tweet) error
}

type tweetValidator struct{}

func NewTweetValidator() ITweetValidator {
	return &tweetValidator{}
}

func (tv *tweetValidator) TweetValidate(tweet model.Tweet) error {
	return validation.ValidateStruct(&tweet,
		validation.Field(
			&tweet.Content,
			validation.Required.Error("tweet content is required"),
			validation.RuneLength(1, 140).Error("tweet limit min 1 max 140 chars"),
		),
	)
}
