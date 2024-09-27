package controllers

import (
	"net/http"

	"github.com/aogallo/go-server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	var users []models.User

	if err := uc.DB.Model(&models.User{}).Preload("Roles").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retriving users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": users})
}
