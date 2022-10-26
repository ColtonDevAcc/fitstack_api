package routine

import "github.com/VooDooStack/FitStackAPI/domain/exercise"

type Routine struct {
	ID          int              `json:"id" db:"routine.id"`
	Title       string           `json:"title" db:"routine.title"`
	Description string           `json:"description" db:"routine.description"`
	ImageUrl    string           `json:"image_url" db:"routine.image_url"`
	Schedule    RoutineSchedule  `json:"schedule" db:"routine.schedule"`
	Workouts    exercise.Workout `json:"workouts" db:""`
}
