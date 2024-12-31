package auth

import (
	"fmt"
	"net/http"

	"github.com/aogallo/go-server/internal/models"
	"github.com/aogallo/go-server/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Login struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type AuthController struct {
	DB *gorm.DB
}

func newAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func (auth *AuthController) Login(context *gin.Context) {
	var login Login

	error := context.ShouldBindJSON(&login)

	if error != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, fmt.Sprintf("Login validation failed!. %s", error.Error()))
		return
	}

	var user models.User

	result := auth.DB.Model(&models.User{}).Preload("Roles").Where("username = ?", login.Username).First(&user)

	if result.Error != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid credentials")
		return
	}

	token, error := utils.GenerateToken(user)

	if error != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}
