package exercise

type Workout struct {
	ID          int            `json:"id" db:"workout.id"`
	Name        string         `json:"name" db:"workout.name"`
	WorkoutSets []*WorkoutSets `json:"workout_sets" db:"workout.sets"`
}
