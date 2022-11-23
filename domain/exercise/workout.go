package exercise

import "github.com/VooDooStack/FitStackAPI/domain/user"

type Workout struct {
	ID          uint          `json:"id" gorm:"primaryKey; unique"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	CreatorID   string        `json:"creator_id"`
	Creator     *user.User    `json:"creator" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	WorkoutSets []WorkoutSets `json:"workout_sets" gorm:"foreignKey:ID"`
}

type WorkoutUsecase interface {
	GetById(uuid string) (*Workout, error)
	CreateWorkout(workout *Workout) error
	GetAll(userId string) ([]*Workout, error)
	UpdateWorkout(workout *Workout) error
	DeleteWorkout(workout *Workout) error
	GetWorkoutSets(id uint) (*[]WorkoutSets, error)
}

type WorkoutRepository interface {
	SelectById(uuid string) (*Workout, error)
	SelectAll(userId string) ([]*Workout, error)
	Insert(workout *Workout) error
	Update(workout *Workout) error
	Delete(Workout *Workout) error
	FetchWorkoutSets(id uint) (*[]WorkoutSets, error)
}
