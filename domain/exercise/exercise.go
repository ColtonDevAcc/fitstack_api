package exercise

type Exercise struct {
	Name              string               `json:"name" db:"name"`
	Description       string               `json:"description" db:"description"`
	Image             *string              `json:"image" db:"image"`
	MetValue          float32              `json:"met_value" db:"met_value"`
	ExerciseType      []*ExerciseType      `db:""`
	ExerciseEquipment []*ExerciseEquipment `db:""`
	MuscleTarget      []*MuscleTarget      `db:""`
}

type ExerciseUsecase interface {
	GetById(uuid string) (*Exercise, error)
	CreateExercise(exercise *Exercise) error
	GetExercise() ([]*Exercise, error)
	UpdateExercise(exercise *Exercise) error
}

type ExerciseRepository interface {
	SelectById(uuid string) (*Exercise, error)
	SelectAll(userId string) ([]*Exercise, error)
	Insert(exercise *Exercise) error
	Update(exercise *Exercise) error
}
