package database

import (
	"os/user"

	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&user.User{}); result.Error != nil {
		return result.Error
	}

	return nil
}
