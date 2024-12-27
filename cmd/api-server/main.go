package main

import (
	"github.com/aogallo/go-server/internal/db"
	"github.com/aogallo/go-server/internal/models"
	"github.com/aogallo/go-server/internal/routes"
)

func main() {
	database := db.ConnectDB(".env")
	defer db.DisconnectDB(database)

	// Migrate the schemas
	database.AutoMigrate(&models.User{}, &models.Role{})

	r := routes.SetupRouter(database)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
