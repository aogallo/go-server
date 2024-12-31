package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(path string) *gorm.DB {
	errorEnv := godotenv.Load(path)
	if errorEnv != nil {
		fmt.Printf("testing %+v", errorEnv)
		panic("Failed to load the configurations")
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")

	db, error := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})

	if error != nil {
		panic("Failed to connect database")
	}

	return db
}

func DisconnectDB(db *gorm.DB) {
	dbSql, error := db.DB()

	if error != nil {
		panic("Failed to kill connection from database")
	}

	dbSql.Close()
}
