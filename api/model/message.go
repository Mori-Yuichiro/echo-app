package model

import "time"

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	RoomId    uint      `json:"room_id" gorm:"not null"`
	Message   string    `json:"message" gorm:"not null; default:''"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Room      Room      `json:"room" gorm:"foreignKey:RoomId; constraint:OnDelete:CASCADE"`
}

type MessageResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	RoomId    uint      `json:"room_id" gorm:"not null"`
	Message   string    `json:"message" gorm:"not null; default:''"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
