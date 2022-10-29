package usecase

import (
	"github.com/VooDooStack/FitStackAPI/domain/exercise"
)

type WorkoutUsecase struct {
	workoutRepo exercise.WorkoutRepository
}

func NewWorkoutUsecase(wr exercise.WorkoutRepository) exercise.WorkoutUsecase {
	return &WorkoutUsecase{workoutRepo: wr}
}

func (w *WorkoutUsecase) GetById(uuid string) (*exercise.Workout, error) {
	//TODO:
	return w.workoutRepo.SelectById(uuid)
}

func (w *WorkoutUsecase) CreateWorkout(workout *exercise.Workout) error {
	err := w.workoutRepo.Insert(workout)
	if err != nil {
		return err
	}

	return nil
}

func (w *WorkoutUsecase) GetAll(userId string) ([]*exercise.Workout, error) {
	//TODO:
	return w.workoutRepo.SelectAll(userId)
}

func (w *WorkoutUsecase) UpdateWorkout(workout *exercise.Workout) error {
	err := w.workoutRepo.Update(workout)
	if err != nil {
		return err
	}

	return nil
}
