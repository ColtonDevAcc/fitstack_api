package repository

import (
	"github.com/VooDooStack/FitStackAPI/domain/exercise"
	"gorm.io/gorm"
)

type workoutRepository struct {
	Database gorm.DB
}

func NewWorkoutRepository(db gorm.DB) exercise.WorkoutRepository {
	return &workoutRepository{db}
}

func (wr *workoutRepository) SelectById(uuid string) (*exercise.Workout, error) {
	workout := &exercise.Workout{}
	err := wr.Database.Where("uuid = ?", uuid).Preload("WorkoutSets").Preload("Creator.Profile").First(&workout).Error
	return workout, err
}

func (wr *workoutRepository) SelectAll(creatorId string) ([]*exercise.Workout, error) {
	workouts := []*exercise.Workout{}
	err := wr.Database.Where("creator_id = ?", creatorId).Preload("WorkoutSets.Exercises").Preload("Creator.Profile").Find(&workouts).Error
	return workouts, err
}

func (wr *workoutRepository) Insert(workout *exercise.Workout) error {
	return wr.Database.Create(&workout).Error
}

func (wr *workoutRepository) Update(workout *exercise.Workout) error {
	return wr.Database.Save(workout).Error
}

func (wr *workoutRepository) Delete(uuid string) error {
	return wr.Database.Where("uuid = ?", uuid).Delete(&exercise.Workout{}).Error
}
