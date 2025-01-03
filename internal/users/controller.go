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

func newUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) GetUsers(context *gin.Context) {
	var users []models.User

	if err := uc.DB.Model(&models.User{}).Preload("Roles").Find(&users).Error; err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Error to retrieve users")
		context.Abort()
		return
	}

	responses := make([]models.UserResponse, len(users))

	for i, user := range users {
		responses[i] = user.ToResponse()
	}

	utils.SuccessResponse(context, http.StatusOK, responses)
}

func (uc *UserController) CreateUser(context *gin.Context) {
	var user models.User

	error := context.ShouldBindJSON(&user)

	if error != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, fmt.Sprintf("User validation failed!. %s", error.Error()))
		context.Abort()
		return
	}

	hashedPassword, error := utils.HasPassword(user.Password)

	if error != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, fmt.Sprintf("User validation failed!. %s", error.Error()))
		context.Abort()
		return
	}

	user.Password = hashedPassword
	result := uc.DB.Create(&user)

	if result.Error != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, fmt.Sprintf("User validation failed!. %s", result.Error.Error()))
		context.Abort()
		return
	}

	utils.SimpleSuccessResponse(context, http.StatusOK)
}

func (uc *UserController) GetUserById(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		utils.ErrorResponse(context, http.StatusBadRequest, "User validation failed!. The ID is not validated")
		context.Abort()
		return
	}

	user, result := getUserById(id, uc)

	if result.Error != nil {
		utils.ErrorResponse(context, http.StatusNotFound, fmt.Sprintf("User validation failed!. %s", result.Error.Error()))
		context.Abort()
		return
	}

	if result.RowsAffected == 0 {
		utils.ErrorResponse(context, http.StatusNotFound, "User validation failed!. User Not Found")
		context.Abort()
		return
	}

	utils.SuccessResponse(context, http.StatusOK, user.ToResponse())
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
		utils.ErrorResponse(context, http.StatusBadRequest, "User validation failed!. The ID is not validated")
		context.Abort()
		return
	}

	if err := context.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, fmt.Sprintf("User validation failed!. %s", err.Error()))
		context.Abort()
		return
	}

	userDB, result := getUserById(id, uc)

	if result.Error != nil {
		utils.ErrorResponse(context, http.StatusNotFound, fmt.Sprintf("User validation failed!. %s", result.Error.Error()))
		context.Abort()
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
		utils.ErrorResponse(context, http.StatusNotFound, fmt.Sprintf("User validation failed!. %s", updatedResult.Error.Error()))
		context.Abort()
		return
	}

	utils.SimpleSuccessResponse(context, http.StatusOK)
}

func (uc *UserController) DeleteUser(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		utils.ErrorResponse(context, http.StatusBadRequest, "User validation failed!. The ID is not validated")
		context.Abort()
		return
	}

	userDB, result := getUserById(id, uc)

	if result.Error != nil {
		utils.ErrorResponse(context, http.StatusNotFound, fmt.Sprintf("User validation failed!. %s", result.Error.Error()))
		context.Abort()
		return
	}

	if result.RowsAffected == 0 {
		utils.ErrorResponse(context, http.StatusNotFound, "User validation failed!. User Not Found")
		context.Abort()
		return
	}

	if error := uc.DB.Select("Roles").Delete(&userDB).Error; error != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, fmt.Sprintf("User validation failed!. %s", error.Error()))
		context.Abort()
		return

	}

	utils.SimpleSuccessResponse(context, http.StatusNoContent)
}
