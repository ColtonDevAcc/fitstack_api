package exercise

type Workout struct {
	ID          int            `json:"id" db:"id"`
	Name        string         `json:"name" db:"name"`
	Publisher   *string        `json:"publisher" db:"routine"`
	WorkoutSets []*WorkoutSets `json:"workout_sets" db:"workout_sets"`
}
