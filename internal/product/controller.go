package product

import (
	"fmt"
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

	response := make([]models.ProductResponse, len(products))

	for i, product := range products {
		response[i] = product.ConvertToResponse()
	}

	utils.SuccessResponse(context, http.StatusOK, response)
}

func (pc *ProductController) CreateProduct(context *gin.Context) {
	var product models.Product

	err := context.ShouldBindJSON(&product)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, fmt.Sprintf("Validation failed!. %s", err.Error()))
		return
	}

	result := pc.DB.Create(&product)

	if result.Error != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, result.Error.Error())
		return
	}

	utils.SuccessResponse(context, http.StatusOK, product.ConvertToResponse())
}

func (pc *ProductController) GetProductById(context *gin.Context) {
	var product models.Product

	id := context.Param("id")

	if id == "" {
		utils.ErrorResponse(context, http.StatusBadRequest, "")
		context.Abort()
		return
	}

	result := pc.DB.First(&product, id)

	if result.Error != nil {
		utils.ErrorResponse(context, http.StatusNotFound, "Product Not Found")
	}

	utils.SuccessResponse(context, http.StatusOK, product.ConvertToResponse())

}