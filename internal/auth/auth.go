package auth

import (
	"net/http"

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

	error := context.ShouldBindJSON(login)

	if error != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"success": false, "message": "Login validation failed!", "error": error.Error()})
	}

}
