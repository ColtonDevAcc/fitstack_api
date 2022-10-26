package program

import (
	"github.com/VooDooStack/FitStackAPI/domain/routine"
	"github.com/VooDooStack/FitStackAPI/domain/user"
	"github.com/google/uuid"
)

type Program struct {
	ID          uuid.UUID          `json:"id" db:"id"`
	Title       string             `json:"title" db:"title"`
	Description string             `json:"description" db:"description"`
	Creator     *user.UserProfile  `json:"creator" db:"creator"  gorm:"foreignKey:id"`
	Routine     []*routine.Routine `json:"workouts" db:"routine"  gorm:"foreignKey:id"`
}

type ProgramUsecase interface {
	GetById(uuid string) (*Program, error)
	GetByCreator(creatorId string) (*Program, error)
	Get(uuid string) (*Program, error)
}

type ProgramRepository interface {
	GetById(uuid string) (*Program, error)
	GetByCreator(creatorId string) (*Program, error)
	Get(uuid string) (*Program, error)
}
