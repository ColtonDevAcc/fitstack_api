package routine

import (
	"time"
)

type RoutineSchedule struct {
	ID        int        `json:"id" db:"id"`
	StartDate *time.Time `json:"start_date" db:"start_date"`
	EndDate   *time.Time `json:"end_date" db:"end_date"`
}
