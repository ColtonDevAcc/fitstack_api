package repository

import (
	"context"
	"fmt"

	"github.com/VooDooStack/FitStackAPI/domain/program"
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
	*
    FROM programs as program
	LEFT JOIN routines as routine
    on program.routine_id = routine.id AND program.creator = $1
	LEFT JOIN workouts as workout
    on workout.id = routine.workout_id;
	`

	rows, err := u.Database.Query(context.Background(), sqlStatement, uuid)
	if err != nil {
		logrus.Error(fmt.Printf("error querying row err: %v", err))
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		program := program.Program{}

		err = rows.Scan(
			&program.ID,
			&program.Title,
			&program.Description,
			&program.Creator,
			&program.RoutineID,
			&program.Routine.ID,
			&program.Routine.Title,
			&program.Routine.Description,
			&program.Routine.ImageUrl,
			&program.Routine.WorkoutID,
			&program.Routine.RoutineScheduleID,
			&program.Routine.Workouts.ID,
			&program.Routine.Workouts.Name,
			&program.Routine.Schedule.ID,
		)
		if err != nil {
			logrus.Error(fmt.Printf("error scanning row err: %v", err))
			return nil, err
		}

		programs = append(programs, &program)
	}
	return programs, nil
}
func (u *programRepository) GetByCreator(creatorId string) (*program.Program, error) {
	//TODO:
	return nil, nil
}
