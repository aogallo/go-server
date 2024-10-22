package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	authController := newAuthController(db)

	// Get user value
	r.POST("/login", authController.Login)
}
