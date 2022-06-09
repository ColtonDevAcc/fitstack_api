package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up database...")

	var connectionString string
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	instanceConnectionName := os.Getenv("INSTANCE_CONNECTION_NAME")
	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")

	if !isSet {
		socketDir = "/cloudsql"
	}

	if os.Getenv("LOCAL") == "" || os.Getenv("LOCAL") == "true" {
		connectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	} else {
		connectionString = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUsername, dbPassword, socketDir, instanceConnectionName, dbName)
	}

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to setup database error:", err)
		return db, err
	}

	return db, nil
}
