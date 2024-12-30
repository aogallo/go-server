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

	publicRoute := r.Group("/api/v1")

	// Auth Routes
	auth.SetupAuthRoutes(publicRoute, db)

	apiV1 := r.Group("/api/v1")

	apiV1.Use(middleware.AuthenticationMiddleware())

	// User Routes
	users.SetupUserRoutes(apiV1, db)

	// Role Routes
	roles.SetupRoleRoutes(apiV1, db)

	return r
}
