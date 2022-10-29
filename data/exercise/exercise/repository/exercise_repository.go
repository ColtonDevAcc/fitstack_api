package repository

import (
	"github.com/VooDooStack/FitStackAPI/domain/exercise"
	"github.com/jackc/pgx/v5/pgxpool"
)

type exerciseRepository struct {
	Database pgxpool.Pool
}

func NewExerciseRepository(db pgxpool.Pool) exercise.ExerciseRepository {
	return &exerciseRepository{Database: db}
}

func (e *exerciseRepository) SelectById(uuid string) (*exercise.Exercise, error) {
	return nil, nil
}

func (e *exerciseRepository) Insert(exercise *exercise.Exercise) error {
	return nil
}

func (e *exerciseRepository) Update(exercise *exercise.Exercise) error {
	return nil
}

func (e *exerciseRepository) SelectAll(userId string) ([]*exercise.Exercise, error) {
	return nil, nil
}
