package usecase

import (
	"github.com/VooDooStack/FitStackAPI/domain/exercise"
)

type ExerciseUsecase struct {
	exerciseRepo exercise.ExerciseRepository
}

func NewExerciseUsecase(er exercise.ExerciseRepository) exercise.ExerciseUsecase {
	return &ExerciseUsecase{exerciseRepo: er}
}

func (e *ExerciseUsecase) GetById(uuid string) (*exercise.Exercise, error) {
	return nil, nil
}

func (e *ExerciseUsecase) CreateExercise(exercise *exercise.Exercise) error {
	return nil
}

func (e *ExerciseUsecase) GetExercise() ([]*exercise.Exercise, error) {
	return nil, nil
}

func (e *ExerciseUsecase) UpdateExercise(exercise *exercise.Exercise) error {
	return nil
}
