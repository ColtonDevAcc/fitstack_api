package routine

import "github.com/VooDooStack/FitStackAPI/domain/exercise"

type Routine struct {
	ID                int              `json:"id" db:"id"`
	Title             string           `json:"title" db:"title"`
	Description       string           `json:"description" db:"description"`
	ImageUrl          string           `json:"image_url" db:"image_url"`
	WorkoutID         int              `json:"workout_id" db:"workout_id"`
	RoutineScheduleID int              `json:"routine_schedule" db:"schedule_id"`
	Schedule          RoutineSchedule  `json:"schedule" db:"schedule"`
	Workouts          exercise.Workout `json:"workouts" db:"workouts"`
}
