package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up database...")

	var connectionString string
	dbUsername := mustGetenv("DB_USERNAME")
	dbPassword := mustGetenv("DB_PASSWORD")
	dbHost := mustGetenv("DB_HOST")
	dbTable := mustGetenv("DB_TABLE")
	dbPort := mustGetenv("DB_PORT")
	dbName := mustGetenv("DB_NAME")
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
