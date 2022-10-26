package routine

import (
	"github.com/VooDooStack/FitStackAPI/domain/exercise"
	"gorm.io/gorm"
)

type Routine struct {
	gorm.Model
	Title       string
	Description string
	ImageUrl    string
	Exercise    string
	Schedule    *RoutineSchedule  `gorm:"foreignKey:id"`
	Workout     *exercise.Workout `gorm:"foreignKey:id"`
}
