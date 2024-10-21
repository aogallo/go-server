package roles

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRolRoutes(r *gin.RouterGroup, db *gorm.DB) {
	rolController := newRolController(db)

	r.GET("/roles", rolController.GetRoles)
	r.GET("/roles/:id", rolController.GetRolById)
	r.POST("/roles", rolController.CreateRol)
	r.DELETE("/roles/:id", rolController.DeleteRol)
	r.PUT("/roles/:id", rolController.UpdateRol)
}
