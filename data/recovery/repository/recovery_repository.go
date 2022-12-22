package repository

import (
	"github.com/VooDooStack/FitStackAPI/domain/muscle"
	"gorm.io/gorm"
)

type recoveryRepository struct {
	Database gorm.DB
}

func NewRecoveryRepository(db gorm.DB) muscle.RecoveryRepository {
	return &recoveryRepository{db}
}

func (ur *recoveryRepository) FetchRecovery(uuid string) (*muscle.Recovery, error) {
	recovery := &muscle.Recovery{}
	err := ur.Database.Model(muscle.Recovery{}).Where("user_id = ?", uuid).Preload("Muscles").Statement.Order("updated_at desc").First(&recovery).Error

	return recovery, err
}
