package exercise

import "github.com/google/uuid"

type Workout struct {
	ID          int
	Name        string
	ProgramId   *uuid.UUID
	Publisher   *string
	WorkoutSets []*WorkoutSets
}
