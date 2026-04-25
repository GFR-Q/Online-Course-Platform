package models

import "time"

type Course struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}
