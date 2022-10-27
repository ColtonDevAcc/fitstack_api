package repository

import (
	"context"
	"fmt"

	"github.com/VooDooStack/FitStackAPI/domain/program"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type programRepository struct {
	Database pgxpool.Pool
}

func NewProgramRepository(db pgxpool.Pool) program.ProgramRepository {
	return &programRepository{db}
}

func (u *programRepository) GetById(uuid string) (*program.Program, error) {
	//TODO:
	return nil, nil
}

func (u *programRepository) Get(uuid string) ([]*program.Program, error) {
	programs := []*program.Program{}

	sqlStatement := `
	SELECT
    program.*,
    routine.id as "routine.id",
    routine.title as "routine.title",
    routine.description as "routine.description",
    routine.image_url as "routine.image_url",
    workout.id as "workout.id",
    workout.name as "workout.name",
    array_to_json(array_agg(workout_set.*)) as "workout.sets"

    FROM programs as program
	LEFT JOIN routines as routine
    on program.routine_id = routine.id AND program.creator = $1
	LEFT JOIN workouts as workout
    on workout.id = routine.workout_id
    LEFT JOIN workout_sets as workout_set
    on workout_set.workout_id = workout.id
    GROUP BY program.id, routine.id, workout.id
	`

	rows, err := u.Database.Query(context.Background(), sqlStatement, uuid)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("error querying row err: %v", err)
	}

	err = pgxscan.ScanAll(&programs, rows)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("error scanning rows: %v", err)
	}

	return programs, nil
}
func (u *programRepository) GetByCreator(creatorId string) (*program.Program, error) {
	//TODO:
	return nil, nil
}
