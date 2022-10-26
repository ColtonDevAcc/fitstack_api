package repository

import (
	"github.com/VooDooStack/FitStackAPI/domain/program"
	"gorm.io/gorm"
)

type programRepository struct {
	Database gorm.DB
}

func NewProgramRepository(db gorm.DB) program.ProgramRepository {
	return &programRepository{db}
}

func (u *programRepository) GetById(uuid string) (*program.Program, error) {
	//TODO:
	return nil, nil
}

func (u *programRepository) Get(uuid string) (*program.Program, error) {
	//TODO:
	programs := []*program.Program{}
	// sqlStatement := `
	// SELECT * FROM programs as "programs"
	// LEFT JOIN workouts as "workouts"
	// on workouts.program_id = programs.id
	// WHERE programs.creator=$1;
	// `

	// rows, err := u.Database.Query(context.Background(), sqlStatement, uuid)
	// if err != nil {
	// 	logrus.Error(fmt.Printf("error querying row err: %v", err))
	// 	return nil, err
	// }

	// err = pgxscan.ScanAll(&programs, rows)
	// if err != nil {
	// 	logrus.Error(err)
	// 	return nil, err
	// }

	return programs[1], nil
}
func (u *programRepository) GetByCreator(creatorId string) (*program.Program, error) {
	//TODO:
	return nil, nil
}
