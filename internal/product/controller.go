package product

import (
	"net/http"

	"github.com/aogallo/go-server/internal/models"
	"github.com/aogallo/go-server/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func newProductController(db *gorm.DB) *ProductController {
	return &ProductController{DB: db}
}

func (pc *ProductController) GetProducts(context *gin.Context) {
	var products []models.Product

	result := pc.DB.Find(&products)

	if result.Error != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Fail to retrieve the products")
		return
	}

	utils.SuccessResponse(context, http.StatusOK, products)
}
