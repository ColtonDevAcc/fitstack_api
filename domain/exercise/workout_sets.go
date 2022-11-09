package exercise

type WorkoutSets struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	Exercises []*Exercise `json:"exercises" gorm:"foreignKey:ID"`
}
