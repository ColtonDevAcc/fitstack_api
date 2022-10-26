package database

import (
	"github.com/VooDooStack/FitStackAPI/domain/exercise"
	"github.com/VooDooStack/FitStackAPI/domain/program"
	"github.com/VooDooStack/FitStackAPI/domain/user"
	"gorm.io/gorm"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&program.Program{}, &user.User{}, &user.Friendship{}, &exercise.Exercise{}, &exercise.ExerciseEquipment{}, &exercise.ExerciseType{}, &exercise.MuscleTarget{}, &exercise.Workout{}, &exercise.WorkoutSets{})
	return err
}
