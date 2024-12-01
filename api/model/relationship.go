package model

import "time"

type Relationship struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	FollowerId uint      `json:"follower_id" gorm:"not null; uniqueIndex:idx_follow_follower"`
	FollowedId uint      `json:"followed_id" gorm:"not null; uniqueIndex:idx_follow_follower"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Follower   User      `json:"follower" gorm:"foreignKey:FollowerId; constraint:OnDelete:CASCADE"`
	Followed   User      `json:"followed" gorm:"foreignKey:FollowedId; constraint:OnDelete:CASCADE"`
}

type RelationshipResponse struct {
	ID         uint         `json:"id" gorm:"primaryKey"`
	FollowerId uint         `json:"follower_id" gorm:"not null; uniqueIndex:idx_follow_follower"`
	FollowedId uint         `json:"followed_id" gorm:"not null; uniqueIndex:idx_follow_follower"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	Follower   UserResponse `json:"follower"`
	Followed   UserResponse `json:"followed"`
}
