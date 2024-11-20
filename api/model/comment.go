package model

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Comment   string    `json:"comment" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	TweetId   uint      `json:"tweet_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Tweet     Tweet     `json:"tweet" gorm:"foreignKey:TweetId; constraint:OnDelete:CASCADE"`
}

type CommentResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Comment   string    `json:"comment" gorm:"not null"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	TweetId   uint      `json:"tweet_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
