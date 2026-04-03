package models

type Course struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	UserID      uint    `json:"user_id"`
}
