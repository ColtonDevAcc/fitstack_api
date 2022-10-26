package routine

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoutineSchedule struct {
	gorm.Model
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Monday    uuid.UUID `json:"monday"`
	Tuesday   uuid.UUID `json:"tuesday"`
	Wednesday uuid.UUID `json:"wednesday"`
	Thursday  uuid.UUID `json:"thursday"`
	Friday    uuid.UUID `json:"friday"`
	Saturday  uuid.UUID `json:"saturday"`
	Sunday    uuid.UUID `json:"sunday"`
}
