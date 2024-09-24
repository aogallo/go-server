package controllers

import (
	"net/http"

	"github.com/aogallo/go-server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RolController struct {
	DB *gorm.DB
}

func NewRolController(db *gorm.DB) *RolController {
	return &RolController{DB: db}
}

func (rc *RolController) GetRoles(c *gin.Context) {
	var roles []models.Rol

	c.JSON(http.StatusOK, roles)
}
