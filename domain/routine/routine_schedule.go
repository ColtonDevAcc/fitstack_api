package routine

import (
	"time"

	"gorm.io/gorm"
)

type RoutineSchedule struct {
	gorm.Model
	StartDate *time.Time `json:"start_date" db:"start_date"`
	EndDate   *time.Time `json:"end_date" db:"end_date"`
}
