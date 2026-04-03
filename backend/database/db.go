package database

import (
	"log"

	"github.com/GFR-Q/Online-Course-Platform/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=127.0.0.1 user=postgres password= dbname=course_db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	DB.AutoMigrate(&models.User{}, &models.Course{})
}
