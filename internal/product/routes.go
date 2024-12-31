package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductRoutes(router *gin.RouterGroup, db *gorm.DB) {
	productController := newProductController(db)

	router.GET("/products", productController.GetProducts)
}
