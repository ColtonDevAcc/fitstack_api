package repository

import (
	"fmt"

	"github.com/VooDooStack/FitStackAPI/domain/exercise"
	"gorm.io/gorm"
)

type exerciseRepository struct {
	Database gorm.DB
}

func NewExerciseRepository(db gorm.DB) exercise.ExerciseRepository {
	return &exerciseRepository{Database: db}
}

func (e *exerciseRepository) SelectById(uuid string) (*exercise.Exercise, error) {
	exercise := &exercise.Exercise{}
	err := e.Database.Where("id = ?", uuid).Preload("ExerciseTypes").Preload("ExerciseEquipment").Preload("MuscleTargets").Preload("Creator.Profile").First(&exercise).Error
	return exercise, err
}

func (e *exerciseRepository) Insert(exercise *exercise.Exercise) error {
	e.Database.Model(&exercise).Association("ExerciseEquipment").Replace(exercise.ExerciseEquipment)
	e.Database.Model(&exercise).Association("MuscleTargets").Replace(exercise.MuscleTarget)
	return e.Database.Create(exercise).Error
}

func (e *exerciseRepository) Update(ex *exercise.Exercise) error {
	err := e.Database.Model(exercise.Exercise{}).Where("id = ?", ex.ID).Updates(&ex).Error
	return err
}

func (e *exerciseRepository) SelectAll(userId string) (*[]exercise.Exercise, error) {
	exercises := []exercise.Exercise{}
	err := e.Database.Model(exercise.Exercise{}).Preload("ExerciseEquipment").Preload("MuscleTarget").Preload("Creator.Profile").Find(&exercises).Error
	fmt.Printf("exercises: %v", &exercises)
	return &exercises, err
}

func (e *exerciseRepository) Delete(exercise *exercise.Exercise) error {
	err := e.Database.Where("id = ?", exercise.ID).Delete(&exercise).Error
	return err
}

func (e *exerciseRepository) FetchExerciseTypes(id uint) (*[]exercise.ExerciseType, error) {
	exerciseTypes := []exercise.ExerciseType{}
	err := e.Database.Model(&exercise.Exercise{}).Where("id = ?", id).First(&exerciseTypes).Error

	return &exerciseTypes, err
}

func (e *exerciseRepository) FetchExerciseEquipment(id uint) (*[]exercise.ExerciseEquipment, error) {
	ExerciseEquipment := []exercise.ExerciseEquipment{}
	err := e.Database.Model(&exercise.ExerciseEquipment{}).Where("id = ?", id).Preload("ExerciseEquipment").First(&ExerciseEquipment).Error

	return &ExerciseEquipment, err
}

func (e *exerciseRepository) FetchMuscleTargets(id uint) (*[]exercise.MuscleTarget, error) {
	MuscleTargets := []exercise.MuscleTarget{}
	err := e.Database.Model(&exercise.MuscleTarget{}).Where("id = ?", id).Preload("MuscleTargets").First(&MuscleTargets).Error

	return &MuscleTargets, err
}
