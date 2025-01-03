package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductRoutes(router *gin.RouterGroup, db *gorm.DB) {
	productController := newProductController(db)

	router.GET("/products", productController.GetProducts)
	router.GET("/products/:id", productController.GetProductById)
	router.POST("/products", productController.CreateProduct)
	router.DELETE("/products/:id", productController.DeleteProductById)
	router.PUT("/products/:id", productController.UpdateProductById)
}
