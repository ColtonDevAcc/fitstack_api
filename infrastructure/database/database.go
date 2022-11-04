package database

import (
	"fmt"

	"github.com/VooDooStack/FitStackAPI/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config config.Config) (*gorm.DB, error) {
	fmt.Println("Setting up database...")

	db, err := gorm.Open(postgres.Open(MustGetenv("DB_URL")), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to setup database error:", err)
		return db, err
	}

	db.Logger.LogMode(logger.Error)

	err = MigrateDB(db)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return db, nil
}

//go.sum,Dockerfile,*.yaml,*.json,.gitignore,.gcloudignore,.env,*.md,Makefile
