package exercise

type WorkoutSets struct {
	ID       int         `json:"id" db:"id"`
	Exercise []*Exercise `json:"exercise" db:"workout_sets.exercises"`
}
