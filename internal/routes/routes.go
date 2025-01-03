package routes

import (
	_ "github.com/aogallo/go-server/docs"
	"github.com/aogallo/go-server/internal/auth"
	"github.com/aogallo/go-server/internal/middleware"
	"github.com/aogallo/go-server/internal/product"
	"github.com/aogallo/go-server/internal/roles"
	"github.com/aogallo/go-server/internal/users"
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

	// url := ginSwagger.URL("http://localhost:8080/swagger/swagger.json")
	publicRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
