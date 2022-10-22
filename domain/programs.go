package domain

import "github.com/google/uuid"

type Program struct {
	ID                uuid.UUID    `json:"id"`
	Title             string       `json:"title"`
	Description       string       `json:"description"`
	Creator           *UserProfile `json:"creator"`
	ExerciseRoutineId uuid.UUID    `json:"exercise_routine_id"`
}

type ProgramUsecase interface {
}

type ProgramRepository interface {
}
