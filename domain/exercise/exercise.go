package exercise

import (
	"github.com/VooDooStack/FitStackAPI/domain/user"
	"github.com/lib/pq"
)

type Exercise struct {
	ID                uint                `json:"id" gorm:"primaryKey"`
	Name              string              `json:"name"`
	Description       string              `json:"description"`
	Images            pq.StringArray      `json:"images" gorm:"type:character varying[]"`
	MetValue          float32             `json:"met_value"`
	CreatorID         string              `json:"creator_id"`
	Creator           *user.User          `json:"creator" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ExerciseEquipment []ExerciseEquipment `json:"exercise_equipment" gorm:"many2many:exercise_equipment;"`
	MuscleTarget      []MuscleTarget      `json:"muscle_targets" gorm:"many2many:exercise_muscle_targets;"`
	Type              *ExerciseType       `json:"type" gorm:"type:exercise_type"`
}

type ExerciseType string

const (
	CARDIO      ExerciseType = "cardio"
	PUSH        ExerciseType = "push"
	PULL        ExerciseType = "pull"
	STATIC      ExerciseType = "static"
	LEGS        ExerciseType = "legs"
	FLEXIBILITY ExerciseType = "flexibility"
	OTHER       ExerciseType = "other"
)

type ExerciseUsecase interface {
	GetById(uuid string) (*Exercise, error)
	CreateExercise(exercise *Exercise) error
	GetExercises(userId string) (*[]Exercise, error)
	UpdateExercise(exercise *Exercise) error
	DeleteExercise(exercise *Exercise) error
	GetExerciseTypes(id uint) (*[]ExerciseType, error)
	GetExerciseEquipment(id uint) (*[]ExerciseEquipment, error)
	GetMuscleTargets(id uint) (*[]MuscleTarget, error)
}

type ExerciseRepository interface {
	SelectById(uuid string) (*Exercise, error)
	SelectAll(userId string) (*[]Exercise, error)
	Insert(exercise *Exercise) error
	Update(exercise *Exercise) error
	Delete(exercise *Exercise) error
	FetchExerciseTypes(id uint) (*[]ExerciseType, error)
	FetchExerciseEquipment(id uint) (*[]ExerciseEquipment, error)
	FetchMuscleTargets(id uint) (*[]MuscleTarget, error)
}
