package program

import (
	"github.com/VooDooStack/FitStackAPI/domain/routine"
	"github.com/VooDooStack/FitStackAPI/domain/user"
	"github.com/google/uuid"
)

type Program struct {
	ID                uuid.UUID          `json:"id"`
	Title             string             `json:"title"`
	Description       string             `json:"description"`
	Creator           *user.UserProfile  `json:"creator"`
	ExerciseRoutineId uuid.UUID          `json:"exercise_routine_id"`
	Routine           []*routine.Routine `json:"workouts" db:"workouts"`
}

type ProgramUsecase interface {
	GetById(uuid string) (*Program, error)
	GetByCreator(creatorId string) (*Program, error)
	Get(uuid string) (*Program, error)
}

type ProgramRepository interface {
	GetById(uuid string) (*Program, error)
	GetByCreator(creatorId string) (*Program, error)
	Get(uuid string) (*Program, error)
}
