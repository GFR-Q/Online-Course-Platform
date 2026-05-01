package main

import (
	"github.com/GFR-Q/Online-Course-Platform/backend/database"
	"github.com/GFR-Q/Online-Course-Platform/backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()
	r.POST("/register", handlers.CreateUser)
	r.POST("/login", handlers.Login)

	r.GET("/courses", handlers.GetCourses)
	r.GET("/courses/:id", handlers.GetCourse)
	r.GET("/search", handlers.SearchCourses)
	r.GET("/courses/convert", handlers.GetConvertedPrice)
	protected := r.Group("/")
	protected.Use(handlers.AuthMiddleware())
	{
		protected.POST("/courses", handlers.CreateCourse)
		protected.PUT("/courses/:id", handlers.UpdateCourse)
		protected.DELETE("/courses/:id", handlers.DeleteCourse)
		protected.GET("/users", handlers.GetUsers)
		protected.GET("/users/:id", handlers.GetUser)
		protected.PUT("/users/:id", handlers.UpdateUser)
		protected.DELETE("/users/:id", handlers.DeleteUser)
	}

	r.Run(":8080")
}
