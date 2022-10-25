package exercise

type WorkoutSets struct {
	ID       int       `json:"int" db:"int"`
	Exercise *Exercise `json:"exercise" db:"exercise"`
}
