package routes

import (
	"github.com/aogallo/go-server/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	// Get user value
	r.GET("/users", userController.GetUsers)
	r.GET("/users/:id", userController.GetUserById)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:id", userController.UpdateUser)
}