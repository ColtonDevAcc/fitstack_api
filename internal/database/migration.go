package database

import (
	"github.com/VooDooStack/FitStackAPI/internal/comment"
	"gorm.io/gorm"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result != nil {
		return result
	}

	return nil
}
