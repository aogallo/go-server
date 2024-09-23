package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	errorEnv := godotenv.Load()
	if errorEnv != nil {
		panic("Failed to load the configurations")
	}

	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	// dbHost := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_NAME")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbName)

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=9920 sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName)

	dsn := "postgresql://biz-ops_owner:n0xM4PSrhjQy@ep-morning-mode-a5k4w3r1-pooler.us-east-2.aws.neon.tech/biz-ops?sslmode=require"

	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		panic("Failed to connect mysql database")
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
