package model

import "time"

type Room struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Entries   []Entry   `json:"entries"`
	Messages  []Message `json:"messages"`
}

type RoomResponse struct {
	ID        uint              `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Entries   []EntryResponse   `json:"entries"`
	Messages  []MessageResponse `json:"messages"`
}
