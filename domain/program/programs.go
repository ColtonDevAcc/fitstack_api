package program

import (
	"github.com/VooDooStack/FitStackAPI/domain/routine"
	"github.com/google/uuid"
)

type Program struct {
	ID          uuid.UUID        `json:"id" db:"id"`
	Title       string           `json:"title" db:"title"`
	Description string           `json:"description" db:"description"`
	Creator     string           `json:"creator" db:"creator"`
	Routine     *routine.Routine `json:"routine" db:""`
}

type ProgramUsecase interface {
	GetById(uuid string) (*Program, error)
	GetByCreator(creatorId string) (*Program, error)
	Get(uuid string) ([]*Program, error)
	Create(program *Program) error
	Update(program *Program) error
}

type ProgramRepository interface {
	GetById(uuid string) (*Program, error)
	GetByCreator(creatorId string) (*Program, error)
	Get(uuid string) ([]*Program, error)
	Create(program *Program) error
	Update(program *Program) error
}
