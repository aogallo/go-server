package routes

import (
	_ "github.com/aogallo/go-server/docs"
	"github.com/aogallo/go-server/internal/auth"
	"github.com/aogallo/go-server/internal/middleware"
	"github.com/aogallo/go-server/internal/v1/product"
	"github.com/aogallo/go-server/internal/v1/roles"
	"github.com/aogallo/go-server/internal/v1/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	ginSwagger "github.com/swaggo/gin-swagger"

	// gin-swagger middleware
	swaggerFiles "github.com/swaggo/files"
)

// swagger embed files

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	endpoint := "/api/v1"

	publicRoutes := router.Group(endpoint)

	// Auth Routes
	auth.SetupAuthRoutes(publicRoutes, db)

	protectedRoutes := router.Group(endpoint)

	protectedRoutes.Use(middleware.AuthenticationMiddleware())

	// User Routes
	users.SetupUserRoutes(protectedRoutes, db)

	// Role Routes
	roles.SetupRoleRoutes(protectedRoutes, db)

	// Product Routes
	product.SetupProductRoutes(protectedRoutes, db)

	router.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))

	return router
}
