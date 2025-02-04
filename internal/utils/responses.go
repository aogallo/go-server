package utils

import (
	"github.com/aogallo/go-server/internal/v1/models"
	"github.com/gin-gonic/gin"
)

func SuccessResponse[T any](context *gin.Context, code int, data T) {
	context.JSON(code, models.APIResponse[T]{
		Success: true,
		Data:    data,
	})
}

func ErrorResponse(context *gin.Context, code int, message string) {
	context.JSON(code, gin.H{
		"success": false,
		"error":   message,
	})
}

func SimpleSuccessResponse(context *gin.Context, code int) {
	context.JSON(code, gin.H{
		"success": true,
	})
}
