package exercise

type Workout struct {
	ID          *int           `json:"id" db:"workout.id"`
	Name        *string        `json:"name" db:"workout.name"`
	WorkoutSets []*WorkoutSets `json:"workout_sets" db:"workout.sets"`
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
