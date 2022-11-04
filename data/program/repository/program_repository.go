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

func (u *programRepository) Get(uuid string) ([]*program.Program, error) {
	programs := []program.Program{}
	err := u.Database.Where("creator_id = $1", uuid).Preload("Creator").Find(&programs).Error

	return nil, err
}

func (u *programRepository) GetByCreator(creatorId string) (*program.Program, error) {
	//TODO:
	return nil, nil
}

func (u *programRepository) Create(program *program.Program) error {
	//TODO:
	return nil
}

func (u *programRepository) Update(program *program.Program) error {
	//TODO:
	return nil
}
