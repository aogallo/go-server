package roles

import (
	"net/http"

	"github.com/aogallo/go-server/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type RolController struct {
	DB *gorm.DB
}

var validate *validator.Validate

func newRolController(db *gorm.DB) *RolController {
	return &RolController{DB: db}
}

func (rc *RolController) GetRoles(c *gin.Context) {
	var roles []models.Rol

	result := rc.DB.Find(&roles)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Rol validation failed!", "error": "Error to retrieve the roles", "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": roles})
}

func (rc *RolController) CreateRol(c *gin.Context) {
	var rol models.Rol

	if err := c.ShouldBindJSON(&rol); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Rol validation failed!", "error": err.Error(), "success": false})
		return
	}

	result := rc.DB.Create(&rol)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "The rol could not be created", "error": result.Error.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": rol})
}

func (rc *RolController) DeleteRol(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Rol validation failed!", "error": "The ID is not valided", "success": false})
		return
	}

	var rolDb models.Rol

	result := rc.DB.First(&rolDb, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Rol validation failed!", "error": "The Rol does not exist", "success": false})
		return
	}

	result = rc.DB.Delete(&models.Rol{}, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "The rol could not be deleted", "error": "", "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (rc *RolController) UpdateRol(c *gin.Context) {
	var rol models.Rol
	var rolDb models.Rol

	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Rol validation failed!", "error": "The ID is not valided", "success": false})
		return
	}

	if err := c.ShouldBindJSON(&rol); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Rol validation failed!", "error": err.Error(), "success": false})
		return
	}

	result := rc.DB.First(&rolDb, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Rol validation failed!", "error": "The Rol does not exist", "success": false})
		return
	}

	updatedResult := rc.DB.Model(&rolDb).Update("name", rol.Name)

	if updatedResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Rol validation failed!", "error": "The Rol can not be updated", "success": false})
		return
	}

	c.JSON(http.StatusOK, rolDb)
}

func (rc *RolController) GetRolById(c *gin.Context) {
	var rol models.Rol
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Rol validation failed!", "error": "The ID is not valided", "success": false})
		return
	}

	result := rc.DB.First(&rol, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Rol validation failed!", "error": "The Rol does not exist", "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": rol})
}
