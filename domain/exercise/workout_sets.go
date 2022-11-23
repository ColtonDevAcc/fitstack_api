package exercise

type WorkoutSets struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	Exercises []*Exercise `json:"exercises" gorm:"many2many:workout_sets_exercises;"`
}
