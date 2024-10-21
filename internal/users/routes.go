package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	userController := NewUserController(db)

	// Get user value
	r.GET("/users", userController.GetUsers)
	r.GET("/users/:id", userController.GetUserById)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
}
