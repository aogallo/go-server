package routes

import (
	"github.com/aogallo/go-server/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {

	userController := controllers.NewUserController(db)

	// Get user value
	r.GET("/users", userController.GetUsers)
}
