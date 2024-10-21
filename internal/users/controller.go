package users

import (
	"fmt"
	"net/http"

	"github.com/aogallo/go-server/internal/models"
	"github.com/aogallo/go-server/internal/utils"
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

	fmt.Printf("usuarios %v", users)

	responses := make([]models.UserResponse, len(users))
	for i, user := range users {
		responses[i] = user.ToResponse()
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": responses})
}

func (uc *UserController) CreateUser(context *gin.Context) {
	var user models.User

	error := context.ShouldBindJSON(&user)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "User validation falied!", "error": error.Error()})
		return
	}

	hashedPassword, error := utils.HasPassword(user.Password)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "User validation falied!", "error": error.Error()})
		return
	}

	user.Password = hashedPassword
	result := uc.DB.Create(&user)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "User validation falied!", "error": result.Error.Error()})
		return
	}

	result = uc.DB.Save(&user)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "User validation falied!", "error": result.Error.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": true})
}

func (uc *UserController) GetUserById(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User validation failed!", "error": "The ID is not valided", "success": false})
		return
	}

	user, result := getUserById(id, uc)

	if result.Error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "User validation failed!", "error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "User validation failed!", "error": "User Not Found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": true, "data": user.ToResponse()})
}

func getUserById(id string, uc *UserController) (models.User, *gorm.DB) {
	var user models.User
	result := uc.DB.Model(&models.User{}).Preload("Roles").First(&user, id)

	return user, result
}

func (uc *UserController) UpdateUser(context *gin.Context) {
	var user models.UserUpdate
	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User validation failed!", "error": "The ID is not valided", "success": false})
		return
	}

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User validation failed!", "error": err.Error(), "success": false})
		return
	}

	userDB, error := getUserById(id, uc)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "User validation failed!", "error": error.Error()})
		return
	}

	updatedResult := uc.DB.Model(&userDB).Updates(&models.User{
		ID:        userDB.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Roles:     user.Roles,
	})

	if updatedResult.Error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User validation failed!",
			"error":   updatedResult.Error.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"success": true, "data": ""})
}

func (uc *UserController) DeleteUser(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User validation failed!", "error": "The ID is not valided", "success": false})
		return
	}

	userDB, result := getUserById(id, uc)

	if result.Error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "User validation failed!", "error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "User validation failed!", "error": "User Not Found"})
		return
	}

	if error := uc.DB.Select("Roles").Delete(&userDB).Error; error != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "User validation failed!", "error": error.Error()})
		return

	}

	context.IndentedJSON(http.StatusNoContent, gin.H{})
}
