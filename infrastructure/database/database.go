package database

import (
	"fmt"

	"github.com/VooDooStack/FitStackAPI/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config config.Config) (*gorm.DB, error) {
	fmt.Println("Setting up database...")

	//!TODO: this should be replaced by config
	dbUsername := MustGetenv("DB_USERNAME")
	dbPassword := MustGetenv("DB_PASSWORD")
	dbHost := MustGetenv("DB_HOST")
	dbName := MustGetenv("DB_NAME")
	dbPort := MustGetenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbName, dbPassword)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to setup database error:", err)
		return db, err
	}

	return db, nil
}

//go.sum,Dockerfile,*.yaml,*.json,.gitignore,.gcloudignore,.env,*.md,Makefile
