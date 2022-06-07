package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up database...")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("failed to setup database error:", err)
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		fmt.Println("failed to ping database error:", err)
		return db, err
	}

	return db, nil
}
