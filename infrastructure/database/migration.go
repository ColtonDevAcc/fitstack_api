package database

import (
	"github.com/VooDooStack/FitStackAPI/models"
	"gorm.io/gorm"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(models.User{})
	if err != nil {
		return err
	}

	return nil
}
