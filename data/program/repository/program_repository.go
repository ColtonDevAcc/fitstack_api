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
	program := &program.Program{}
	err := u.Database.Where("id = $1", uuid).Preload("Creator").First(&program).Error

	return program, err
}

func (u *programRepository) Get(uuid string) ([]*program.Program, error) {
	programs := []*program.Program{}
	err := u.Database.Where("creator_id = $1", uuid).Preload("Creator.Profile").Find(&programs).Error

	return programs, err
}

func (u *programRepository) GetByCreator(creatorId string) (*program.Program, error) {
	program := &program.Program{}
	err := u.Database.Where("creator_id = $1", creatorId).Preload("Creator.Profile").First(&program).Error

	return program, err
}

func (u *programRepository) Create(program *program.Program) error {
	err := u.Database.FirstOrCreate(&program).Error
	return err
}

func (u *programRepository) Update(program *program.Program) error {
	err := u.Database.Save(&program).Error
	return err
}

func (u *programRepository) Delete(id uint, creatorId string) error {
	err := u.Database.Delete(&program.Program{ID: id, CreatorID: creatorId}).Error
	return err
}
