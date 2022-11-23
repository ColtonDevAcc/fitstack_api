package usecase

import (
	"fmt"

	"github.com/VooDooStack/FitStackAPI/domain/exercise"
)

type ExerciseUsecase struct {
	exerciseRepo exercise.ExerciseRepository
}

func NewExerciseUsecase(er exercise.ExerciseRepository) exercise.ExerciseUsecase {
	return &ExerciseUsecase{exerciseRepo: er}
}

func (e *ExerciseUsecase) GetById(uuid string) (*exercise.Exercise, error) {
	ex, err := e.exerciseRepo.SelectById(uuid)
	return ex, err
}

func (e *ExerciseUsecase) CreateExercise(exercise *exercise.Exercise) error {
	err := e.exerciseRepo.Insert(exercise)
	return err
}

func (e *ExerciseUsecase) GetExercises(userId string) (*[]exercise.Exercise, error) {
	fmt.Println("GetExercises")
	ex, err := e.exerciseRepo.SelectAll(userId)
	return ex, err
}

func (e *ExerciseUsecase) UpdateExercise(exercise *exercise.Exercise) error {
	err := e.exerciseRepo.Update(exercise)
	return err
}

func (e *ExerciseUsecase) DeleteExercise(exercise *exercise.Exercise) error {
	err := e.exerciseRepo.Delete(exercise)
	return err
}

func (e *ExerciseUsecase) GetExerciseTypes(id uint) (*[]exercise.ExerciseType, error) {
	et, err := e.exerciseRepo.FetchExerciseTypes(id)
	return et, err
}
func (e *ExerciseUsecase) GetExerciseEquipment(id uint) (*[]exercise.ExerciseEquipment, error) {
	ee, err := e.exerciseRepo.FetchExerciseEquipment(id)
	return ee, err
}
func (e *ExerciseUsecase) GetMuscleTargets(id uint) (*[]exercise.MuscleTarget, error) {
	mt, err := e.exerciseRepo.FetchMuscleTargets(id)
	return mt, err
}
