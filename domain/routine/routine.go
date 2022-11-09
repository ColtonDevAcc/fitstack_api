package routine

import (
	"github.com/VooDooStack/FitStackAPI/domain/exercise"
	"gorm.io/gorm"
)

type Routine struct {
	gorm.Model
	Title       string              `json:"title"`
	Description string              `json:"description"`
	ImageUrl    string              `json:"image_url"`
	Schedule    *RoutineSchedule    `json:"schedule" gorm:"foreignKey:ID"`
	Workouts    []*exercise.Workout `json:"workouts" gorm:"many2many:routine_workouts"`
}
