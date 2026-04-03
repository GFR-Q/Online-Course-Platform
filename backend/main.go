package main

import (
	"github.com/GFR-Q/Online-Course-Platform/backend/database"
	"github.com/GFR-Q/Online-Course-Platform/backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.POST("/courses", handlers.CreateCourse)
	r.GET("/courses", handlers.GetCourses)
	r.GET("/courses/:id", handlers.GetCourse)
	r.PUT("/courses/:id", handlers.UpdateCourse)
	r.DELETE("/courses/:id", handlers.DeleteCourse)
	r.GET("/search", handlers.SearchCourses)

	r.Run(":8080")
}
