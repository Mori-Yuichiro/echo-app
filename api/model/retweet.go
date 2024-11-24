package model

import "time"

type Retweet struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id" gorm:"not null; uniqueIndex:idx_user_tweet"`
	TweetId   uint      `json:"tweet_id" gorm:"not null; uniqueIndex:idx_user_tweet"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Tweet     Tweet     `json:"tweet" gorm:"foreignKey:TweetId; constraint:OnDelete:CASCADE"`
}

type RetweetResponse struct {
	ID        uint          `json:"id" gorm:"primaryKey"`
	UserId    uint          `json:"user_id" gorm:"not null; uniqueIndex:idx_user_tweet"`
	TweetId   uint          `json:"tweet_id" gorm:"not null; uniqueIndex:idx_user_tweet"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Tweet     TweetResponse `json:"tweet"`
}
