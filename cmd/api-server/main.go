package main

import (
	"github.com/aogallo/go-server/internal/db"
	"github.com/aogallo/go-server/internal/routes"
	"github.com/aogallo/go-server/internal/v1/models"
)

//	@title			Golang Rest API
//	@version		1.0
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Allan Gallo
//	@contact.email	allan.gallo.guerra@gmail.com

//	@host		localhost:8080
//	@BasePath	/api/v1

// @Schemes	http https
func main() {
	database := db.ConnectDB(".env")
	defer db.DisconnectDB(database)

	// Migrate the schemas
	database.AutoMigrate(&models.User{}, &models.Role{}, models.Product{})

	r := routes.SetupRouter(database)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
