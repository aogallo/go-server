package roles

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoleRoutes(router *gin.RouterGroup, db *gorm.DB) {
	roleController := newRoleController(db)

	router.GET("/roles", roleController.GetRoles)
	router.GET("/roles/:id", roleController.GetRoleById)
	router.POST("/roles", roleController.CreateRole)
	router.DELETE("/roles/:id", roleController.DeleteRole)
	router.PUT("/roles/:id", roleController.UpdateRole)
}
