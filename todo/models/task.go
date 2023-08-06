package models

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Title     string    `json:"title"`
	CreatorID int       `json:"creator_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
