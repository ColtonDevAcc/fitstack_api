package database

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Info("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up database...")

	var connectionString string
	dbUsername := mustGetenv("DB_USERNAME")
	dbPassword := mustGetenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if os.Getenv("LOCAL") == "" || os.Getenv("LOCAL") == "true" {
		connectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	} else {
		connectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s", dbHost, dbUsername, dbPassword, dbName)

	}

	log.Info("connecting with the following connection string %s", connectionString)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to setup database error:", err)
		return db, err
	}

	return db, nil
}
