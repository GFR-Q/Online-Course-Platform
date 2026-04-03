package models

type Course struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Price       float64
	UserID      uint
}
