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
	//TODO:
	return nil, nil
}
func (p *programUsecase) GetByCreator(creatorId string) (*program.Program, error) {
	//TODO:
	return nil, nil
}
func (p *programUsecase) Get(uuid string) (*program.Program, error) {
	program, err := p.programRepo.Get(uuid)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return program, nil
}
