package routes

import (
	"github.com/aogallo/go-server/internal/auth"
	"github.com/aogallo/go-server/internal/middleware"
	"github.com/aogallo/go-server/internal/roles"
	"github.com/aogallo/go-server/internal/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	var endpoint = "/api/v1"

	publicRoutes := r.Group(endpoint)

	// Auth Routes
	auth.SetupAuthRoutes(publicRoutes, db)

	protectedRoutes := r.Group(endpoint)

	protectedRoutes.Use(middleware.AuthenticationMiddleware())

	// User Routes
	users.SetupUserRoutes(protectedRoutes, db)

	// Role Routes
	roles.SetupRoleRoutes(protectedRoutes, db)

	return r
}
