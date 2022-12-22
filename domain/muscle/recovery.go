package muscle

import (
	"time"

	"github.com/VooDooStack/FitStackAPI/domain/user"
	"gorm.io/gorm"
)

type Recovery struct {
	UserID    string         `json:"id" gorm:"primaryKey; unique"`
	User      user.User      `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Muscles   []Muscle       `json:"muscles" gorm:"many2many:recovery_muscles;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type RecoveryUsecase interface {
	FetchRecovery(uuid string) (*Recovery, error)
}

type RecoveryRepository interface {
	FetchRecovery(uuid string) (*Recovery, error)
}
