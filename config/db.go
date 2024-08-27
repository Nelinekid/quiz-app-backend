package config

import (
	"fmt"

	"github.com/Gambi18/Quizzo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to PostgreSQL database and return the connection object
func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:2323@localhost:5432/quizzo"), &gorm.Config{})
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
	dsnWithDB := fmt.Sprintf("host=localhost user=postgres password=2323 dbname=%s sslmode=disable", dbName)
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
