package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string   `json:"name"`
	Email    string   `gorm:"unique" json:"email"`
	Password string   `json:"password"`
	Courses  []Course `gorm:"foreignKey:UserID" json:"courses"`
}
