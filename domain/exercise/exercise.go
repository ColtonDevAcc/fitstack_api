package exercise

import "github.com/VooDooStack/FitStackAPI/domain/user"

type Exercise struct {
	ID                uint                `json:"id" gorm:"primaryKey"`
	Name              string              `json:"name"`
	Description       string              `json:"description"`
	Image             string              `json:"image"`
	MetValue          float32             `json:"met_value"`
	CreatorID         string              `json:"creator_id"`
	Creator           *user.User          `json:"creator" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ExerciseType      []ExerciseType      `json:"exercise_types" gorm:"foreignKey:ID"`
	ExerciseEquipment []ExerciseEquipment `json:"exercise_equipment" gorm:"foreignKey:ID"`
	MuscleTarget      []MuscleTarget      `json:"muscle_targets" gorm:"foreignKey:ID"`
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
