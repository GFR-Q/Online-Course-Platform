package main

import (
	"net/http"

	"github.com/GFR-Q/Online-Course-Platform/backend/database"
	"github.com/GFR-Q/Online-Course-Platform/backend/models"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		var user models.User
		c.ShouldBindJSON(&user)
		database.DB.Create(&user)
		c.JSON(http.StatusOK, user)
	})

	r.GET("/users", func(c *gin.Context) {
		var users []models.User
		database.DB.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		var user models.User
		database.DB.First(&user, c.Param("id"))
		c.JSON(http.StatusOK, user)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		var user models.User
		database.DB.Delete(&user, c.Param("id"))
		c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
	})

	r.POST("/courses", func(c *gin.Context) {
		var course models.Course
		c.ShouldBindJSON(&course)
		database.DB.Create(&course)
		c.JSON(http.StatusOK, course)
	})

	r.GET("/courses", func(c *gin.Context) {
		var courses []models.Course
		database.DB.Find(&courses)
		c.JSON(http.StatusOK, courses)
	})

	r.PUT("/courses/:id", func(c *gin.Context) {
		var course models.Course
		database.DB.First(&course, c.Param("id"))
		c.ShouldBindJSON(&course)
		database.DB.Save(&course)
		c.JSON(http.StatusOK, course)
	})

	r.DELETE("/courses/:id", func(c *gin.Context) {
		database.DB.Delete(&models.Course{}, c.Param("id"))
		c.JSON(http.StatusOK, gin.H{"message": "Course Deleted"})
	})

	r.GET("/courses/expensive", func(c *gin.Context) {
		var courses []models.Course
		database.DB.Where("price > ?", 100).Find(&courses)
		c.JSON(http.StatusOK, courses)
	})

	r.GET("/courses/sorted", func(c *gin.Context) {
		var courses []models.Course
		database.DB.Order("price asc").Find(&courses)
		c.JSON(http.StatusOK, courses)
	})

	r.Run(":8080")
}
