package usecase

import (
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"

	"github.com/VooDooStack/FitStackAPI/domain/program"
	"github.com/sirupsen/logrus"
)

type programUsecase struct {
	programRepo program.ProgramRepository
	client      auth.Client
	storage     *storage.Client
}

func NewProgramUseCase(pr program.ProgramRepository, client auth.Client, storage *storage.Client) program.ProgramUsecase {
	return &programUsecase{programRepo: pr, client: client, storage: storage}
}

func (p *programUsecase) GetById(uuid string) (*program.Program, error) {
	program, err := p.programRepo.GetById(uuid)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return program, nil
}

func (p *programUsecase) GetByCreator(creatorId string) (*program.Program, error) {
	program, err := p.programRepo.GetByCreator(creatorId)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return program, nil
}

func (p *programUsecase) Get(uuid string) ([]*program.Program, error) {
	program, err := p.programRepo.Get(uuid)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return program, nil
}

func (p *programUsecase) Create(program *program.Program) error {
	err := p.programRepo.Create(program)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *programUsecase) Update(program *program.Program) error {
	err := p.programRepo.Update(program)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *programUsecase) Delete(id uint, creatorId string) error {
	err := p.programRepo.Delete(id, creatorId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
