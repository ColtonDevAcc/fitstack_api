package database

import (
	"github.com/VooDooStack/FitStackAPI/domain"
	"gorm.io/gorm"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(domain.User{}, domain.Friendship{})
	if err != nil {
		return err
	}

	return nil
}
