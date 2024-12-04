package model

import "time"

type Entry struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id" gorm:"not null; uniqueIndex:idx_user_entry"`
	RoomId    uint      `json:"room_id" gorm:"not null; uniqueIndex:idx_user_entry"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Room      Room      `json:"room" gorm:"foreignKey:RoomId; constraint:OnDelete:CASCADE"`
}

type EntryResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	RoomId    uint      `json:"room_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
