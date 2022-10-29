package repository

import (
	"context"

	"github.com/VooDooStack/FitStackAPI/domain/exercise"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type workoutRepository struct {
	Database pgxpool.Pool
}

func NewWorkoutRepository(db pgxpool.Pool) exercise.WorkoutRepository {
	return &workoutRepository{Database: db}
}

func (w *workoutRepository) SelectById(uuid string) (*exercise.Workout, error) {
	var workout exercise.Workout
	sqlStatement := `
	SELECT * FROM workout WHERE id = $1
	`

	err := w.Database.QueryRow(context.Background(), sqlStatement, uuid).Scan(&workout.ID, &workout.Name, &workout.WorkoutSets)
	if err != nil {
		return nil, err
	}

	return &workout, nil
}

func (w *workoutRepository) SelectAll(userId string) ([]*exercise.Workout, error) {
	var workouts []*exercise.Workout
	sqlStatement := `
	SELECT * FROM workout WHERE user_id = $1
	`
	rows, err := w.Database.Query(context.Background(), sqlStatement, userId)
	if err != nil {
		return nil, err
	}

	err = pgxscan.ScanAll(&workouts, rows)
	if err != nil {
		return nil, err
	}

	return workouts, nil
}

func (w *workoutRepository) Insert(workout *exercise.Workout) error {
	sqlStatement := `
	INSERT INTO workout (name) VALUES ($1)
	`

	_, err := w.Database.Exec(context.Background(), sqlStatement, workout.Name)
	if err != nil {
		return err
	}

	return nil
}

func (w *workoutRepository) Update(workout *exercise.Workout) error {
	sqlStatement := `
	UPDATE workout SET name = $1 WHERE id = $2
	`

	_, err := w.Database.Exec(context.Background(), sqlStatement, workout.Name, workout.ID)
	if err != nil {
		return err
	}

	return nil
}
