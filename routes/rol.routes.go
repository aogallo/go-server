package routes

import (
	"github.com/aogallo/go-server/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRolRoutes(r *gin.Engine, db *gorm.DB) {

	rolController := controllers.NewRolController(db)

	r.GET("/rol", rolController.GetRoles)
	r.GET("/rol/:id", rolController.GetRolById)
	r.POST("/rol", rolController.CreateRol)
	r.DELETE("/rol/:id", rolController.DeleteRol)
	r.PUT("/rol/:id", rolController.UpdateRol)

}
