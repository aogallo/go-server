package controllers

import (
	"net/http"

	"github.com/aogallo/go-server/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type RolController struct {
	DB *gorm.DB
}

var validate *validator.Validate

func NewRolController(db *gorm.DB) *RolController {
	return &RolController{DB: db}
}

func (rc *RolController) GetRoles(c *gin.Context) {
	var roles []models.Rol

	c.JSON(http.StatusOK, roles)
}

func (rc *RolController) CreateRol(c *gin.Context) {
	var rol models.Rol

	if err := c.ShouldBindJSON(&rol); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Rol validation failed!", "error": err.Error(), "success": false})
		return
	}

	result := rc.DB.Create(&rol)

	print("result", result, &rol)

	c.JSON(http.StatusOK, rol)
}

func RolStructLevelValidation(sl validator.StructLevel) {
	rol := sl.Current().Interface().(models.Rol)

	if rol.Name == "" {
		sl.ReportError(rol.Name, "Name", "name", "name", "")
	}
}
