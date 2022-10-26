package exercise

import (
	"github.com/VooDooStack/FitStackAPI/domain/user"
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	Name        string           `json:"name"`
	Publisher   user.UserProfile `json:"publisher" gorm:"foreignKey:id"`
	WorkoutSets []WorkoutSets    `json:"workout_sets" gorm:"foreignKey:id"`
}
