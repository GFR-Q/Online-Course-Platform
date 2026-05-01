package handlers

import (
	"net/http"

	"github.com/GFR-Q/Online-Course-Platform/backend/database"
	"github.com/GFR-Q/Online-Course-Platform/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&course)
	c.JSON(http.StatusCreated, course)
}

func GetCourses(c *gin.Context) {
	var courses []models.Course
	database.DB.Find(&courses)
	c.JSON(http.StatusOK, courses)
}

func GetCourse(c *gin.Context) {
	var course models.Course
	if err := database.DB.First(&course, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, course)
}

func UpdateCourse(c *gin.Context) {
	var course models.Course
	if err := database.DB.First(&course, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.ShouldBindJSON(&course)
	database.DB.Save(&course)
	c.JSON(http.StatusOK, course)
}

func DeleteCourse(c *gin.Context) {
	database.DB.Delete(&models.Course{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Course deleted"})
}

func SearchCourses(c *gin.Context) {
	title := c.Query("title")
	var courses []models.Course
	database.DB.Where("title ILIKE ?", "%"+title+"%").Find(&courses)
	c.JSON(http.StatusOK, courses)
}
func GetCoursePriceInUSD(c *gin.Context) {
	price := c.Query("price")
	client := resty.New()

	resp, err := client.R().
		SetQueryParam("price", price).
		SetResult(map[string]interface{}{}).
		Get("http://localhost:8081/convert")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Второй сервис недоступен"})
		return
	}
	c.Data(http.StatusOK, "application/json", resp.Body())
}
