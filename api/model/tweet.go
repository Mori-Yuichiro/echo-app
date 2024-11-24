package model

import "time"

type Tweet struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Content   string     `json:"content" gorm:"not null"`
	ImageUrls string     `json:"image_urls"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	User      User       `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint       `json:"user_id" gorm:"not null"`
	Favorites []Favorite `json:"favorites"`
	Comments  []Comment  `json:"comments"`
	Retweets  []Retweet  `json:"retweets"`
}

type TweetResponse struct {
	ID        uint               `json:"id" gorm:"primaryKey"`
	Content   string             `json:"content" gorm:"not null"`
	ImageUrls []string           `json:"image_urls"`
	User      User               `json:"user" gorm:"not null"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Favorites []FavoriteResponse `json:"favorites"`
	Comments  []CommentResponse  `json:"comments"`
	Retweets  []RetweetResponse  `json:"retweets"`
}
