package roles

import (
	"fmt"
	"net/http"

	"github.com/aogallo/go-server/internal/utils"
	"github.com/aogallo/go-server/internal/v1/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type RoleController struct {
	DB *gorm.DB
}

var validate *validator.Validate

func newRoleController(db *gorm.DB) *RoleController {
	return &RoleController{DB: db}
}

func (rc *RoleController) GetRoles(c *gin.Context) {
	var roles []models.Role

	result := rc.DB.Find(&roles)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Role validation failed!.Error to retrieve the roles")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, roles)
}

func (rc *RoleController) CreateRole(c *gin.Context) {
	var role models.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Role validation failed!. %s", err.Error()))
		return
	}

	result := rc.DB.Create(&role)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("The Role could not be created!. %s", result.Error.Error()))
		return
	}

	utils.SuccessResponse(c, http.StatusOK, role)
}

func (rc *RoleController) DeleteRole(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Role validation failed!. The ID is not validated")
		return
	}

	var roleDb models.Role

	result := rc.DB.First(&roleDb, id)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Role validation failed!. The Role does not exist")
		return
	}

	result = rc.DB.Delete(&models.Role{}, id)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "The Role could not be deleted.")
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
	utils.SimpleSuccessResponse(c, http.StatusOK)
}

func (rc *RoleController) UpdateRole(c *gin.Context) {
	var role models.Role
	var roleDb models.Role

	id := c.Param("id")

	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Role validation failed!. The ID is not validated")
		return
	}

	if err := c.ShouldBindJSON(&role); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Role validation failed!. %s", err.Error()))
		return
	}

	result := rc.DB.First(&roleDb, id)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Role validation failed!. The Role does not exist")
		return
	}

	updatedResult := rc.DB.Model(&roleDb).Update("name", role.Name)

	if updatedResult.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Role validation failed!. The Role can not be updated")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, roleDb)
}

func (rc *RoleController) GetRoleById(c *gin.Context) {
	var role models.Role
	id := c.Param("id")

	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Role validation failed!. The ID is not validated")
		return
	}

	result := rc.DB.First(&role, id)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Role validation failed!. The Role does not exist")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, role)
}
