package database

import (
	"fmt"
	"os"

	"github.com/VooDooStack/FitStackAPI/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config config.Config) (*gorm.DB, error) {
	fmt.Println("Setting up database...")

	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to setup database error:", err)
		return db, err
	}

	err = MigrateDB(db)
	if err != nil {
		fmt.Println(err)
	}

	return db, nil
}

//go.sum,Dockerfile,*.yaml,*.json,.gitignore,.gcloudignore,.env,*.md,Makefile
