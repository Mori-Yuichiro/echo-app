package model

import "time"

type User struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Email           string         `json:"email" gorm:"unique"`
	Password        string         `json:"password"`
	Name            string         `json:"name"`
	Image           string         `json:"image"`
	DisplayName     string         `json:"display_name"`
	PhoneNumber     string         `json:"phone_number"`
	Bio             string         `json:"bio"`
	Location        string         `json:"location"`
	Website         string         `json:"website"`
	Birthday        string         `json:"birthday"`
	ProfileImageUrl string         `json:"profile_image_url"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	Tweets          []Tweet        `json:"tweets"`
	Favorites       []Favorite     `json:"favorites"`
	Comments        []Comment      `json:"comments"`
	Retweets        []Retweet      `json:"retweets"`
	Followers       []Relationship `json:"followers" gorm:"foreignKey:FollowedId; references:ID"`
	Followeds       []Relationship `json:"followeds" gorm:"foreignKey:FollowerId; references:ID"`
}

type UserResponse struct {
	ID              uint                   `json:"id" gorm:"primaryKey"`
	Email           string                 `json:"email" gorm:"unique"`
	Name            string                 `json:"name"`
	Image           string                 `json:"image"`
	DisplayName     string                 `json:"display_name"`
	PhoneNumber     string                 `json:"phone_number"`
	Bio             string                 `json:"bio"`
	Location        string                 `json:"location"`
	Website         string                 `json:"website"`
	Birthday        string                 `json:"birthday"`
	ProfileImageUrl string                 `json:"profile_image_url"`
	Tweets          []TweetResponse        `json:"tweets"`
	Favorites       []FavoriteResponse     `json:"favorites"`
	Comments        []CommentResponse      `json:"comments"`
	Retweets        []RetweetResponse      `json:"retweets"`
	Followers       []RelationshipResponse `json:"followers"`
	Followeds       []RelationshipResponse `json:"followeds"`
}
