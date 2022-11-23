package usecase

import (
	"github.com/VooDooStack/FitStackAPI/domain/exercise"
)

type workoutUsecase struct {
	workoutRepo exercise.WorkoutRepository
}

func NewWorkoutUseCase(wr exercise.WorkoutRepository) exercise.WorkoutUsecase {
	return &workoutUsecase{workoutRepo: wr}
}

func (wuc *workoutUsecase) GetById(uuid string) (*exercise.Workout, error) {
	return wuc.workoutRepo.SelectById(uuid)
}
func (wuc *workoutUsecase) CreateWorkout(workout *exercise.Workout) error {
	return wuc.workoutRepo.Insert(workout)
}
func (wuc *workoutUsecase) GetAll(creatorId string) ([]*exercise.Workout, error) {
	return wuc.workoutRepo.SelectAll(creatorId)
}
func (wuc *workoutUsecase) UpdateWorkout(workout *exercise.Workout) error {
	return wuc.workoutRepo.Update(workout)
}

func (wuc *workoutUsecase) DeleteWorkout(workout *exercise.Workout) error {
	return wuc.workoutRepo.Delete(workout)
}

func (wuc *workoutUsecase) GetWorkoutSets(id uint) (*[]exercise.WorkoutSets, error) {
	return wuc.workoutRepo.FetchWorkoutSets(id)
}
