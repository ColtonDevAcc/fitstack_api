package exercise

type WorkoutSets struct {
	ID       uint        `json:"id" gorm:"primaryKey"`
	Exercise []*Exercise `json:"exercises" gorm:"foreignKey:ID"`
}
