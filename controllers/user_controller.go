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

func (uc *UserController) CreateUser(context *gin.Context) {
	var user models.User

	error := context.ShouldBindJSON(&user)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "User validation falied!", "error": error.Error()})
		return
	}

	result := uc.DB.Create(&user)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "User validation falied!", "error": result.Error.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": true})
}

func (uc *UserController) GetUserById(context *gin.Context) {
	id := context.Param("id")

	user, error := getUserById(id, uc)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "User validation failed!", "error": error.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"success": true, "data": user})
}

func getUserById(id string, uc *UserController) (models.User, error) {
	var user models.User
	result := uc.DB.First(&user, id)

	return user, result.Error
}
