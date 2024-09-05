package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Gambi18/Quizzo/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to PostgreSQL database and return the connection object
func Connect() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	//define env
	DB_URL := os.Getenv("DB_URL")
	DB_nURL := os.Getenv("DB_nURL")

	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database server:", err)
		return
	}

	// Create the database if it doesn't exist
	dbName := "quizzo"
	createDBSQL := fmt.Sprintf(`
    DO $$
    BEGIN
        IF NOT EXISTS (SELECT FROM pg_database WHERE datname = '%s') THEN
            EXECUTE 'CREATE DATABASE %s';
        END IF;
    END
    $$;`, dbName, dbName)

	if err := db.Exec(createDBSQL).Error; err != nil {
		fmt.Println("Failed to create database:", err)
		return
	}

	// Connect to the new database
	dsnWithDB := fmt.Sprintf(DB_nURL, dbName)
	db, err = gorm.Open(postgres.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the specific database:", err)
		return
	}

	// Enable uuid-ossp extension
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		fmt.Println("Failed to enable uuid-ossp extension:", err)
		return
	}

	DB = db
	db.AutoMigrate(&models.User{})
}
