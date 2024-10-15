package main

import (
	"github.com/aogallo/go-server/config"
	"github.com/aogallo/go-server/models"
	"github.com/aogallo/go-server/routes"
)

var db = make(map[string]string)

func main() {
	database := config.ConnectDB(".env")
	defer config.DisconnectDB(database)

	// Migrate the schemas
	database.AutoMigrate(&models.User{}, &models.Rol{})

	r := routes.SetupRouter(database)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
