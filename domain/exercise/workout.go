package exercise

type Workout struct {
	ID          int            `json:"id" db:"workout_id"`
	Name        string         `json:"name"`
	WorkoutSets []*WorkoutSets `json:"workout_sets"`
}
