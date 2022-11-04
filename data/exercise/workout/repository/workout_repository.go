package repository

import (
	"github.com/VooDooStack/FitStackAPI/domain/exercise"
	"gorm.io/gorm"
)

type workoutRepository struct {
	Database gorm.DB
}

func NewWorkoutRepository(db gorm.DB) exercise.WorkoutRepository {
	return &workoutRepository{Database: db}
}

func (w *workoutRepository) SelectById(uuid string) (*exercise.Workout, error) {
	//TODO:

	return nil, nil
}

func (w *workoutRepository) SelectAll(userId string) ([]*exercise.Workout, error) {
	//TODO:

	return nil, nil
}

func (w *workoutRepository) Insert(workout *exercise.Workout) error {
	//TODO:

	return nil
}

func (w *workoutRepository) Update(workout *exercise.Workout) error {
	//TODO:
	return nil
}
