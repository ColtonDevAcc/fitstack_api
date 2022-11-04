package exercise

type Workout struct {
	ID          uint   `json:"id" gorm:"primaryKey; unique"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// Creator     user.User     `json:"creator" gorm:"foreignKey:ID;references:Creator"`
	WorkoutSets []WorkoutSets `json:"workout_sets" gorm:"foreignKey:ID"`
}

type WorkoutUsecase interface {
	GetById(uuid string) (*Workout, error)
	CreateWorkout(*Workout) error
	GetAll(userId string) ([]*Workout, error)
	UpdateWorkout(*Workout) error
}

type WorkoutRepository interface {
	SelectById(uuid string) (*Workout, error)
	SelectAll(userId string) ([]*Workout, error)
	Insert(workout *Workout) error
	Update(workout *Workout) error
}
