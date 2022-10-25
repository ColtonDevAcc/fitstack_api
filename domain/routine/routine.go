package routine

import "github.com/VooDooStack/FitStackAPI/domain/exercise"

type Routine struct {
	Title       string
	Description string
	ImageUrl    string
	Exercise    string
	Schedule    *RoutineSchedule
	Workout     *exercise.Workout
}
