package routes

import (
	"net/http"

	"github.com/aogallo/go-server/internal/auth"
	"github.com/aogallo/go-server/internal/roles"
	"github.com/aogallo/go-server/internal/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	apiV1 := r.Group("/api/v1")

	// Auth Routes
	auth.SetupAuthRoutes(apiV1, db)

	// User Routes
	users.SetupUserRoutes(apiV1, db)

	// Role Routes
	roles.SetupRoleRoutes(apiV1, db)

	return r
}
