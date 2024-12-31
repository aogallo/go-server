package roles

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoleRoutes(r *gin.RouterGroup, db *gorm.DB) {
	rolController := newRolController(db)

	r.GET("/roles", rolController.GetRoles)
	r.GET("/roles/:id", rolController.GetRoleById)
	r.POST("/roles", rolController.CreateRole)
	r.DELETE("/roles/:id", rolController.DeleteRole)
	r.PUT("/roles/:id", rolController.UpdateRole)
}
