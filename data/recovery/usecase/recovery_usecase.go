package usecase

import (
	"github.com/VooDooStack/FitStackAPI/domain/muscle"
)

type userUsecase struct {
	recoveryRepo muscle.RecoveryUsecase
}

func NewRecoveryUseCase(ruc muscle.RecoveryUsecase) muscle.RecoveryUsecase {
	return &userUsecase{recoveryRepo: ruc}
}

func (uu *userUsecase) FetchRecovery(uuid string) (*muscle.Recovery, error) {
	recovery, err := uu.recoveryRepo.FetchRecovery(uuid)

	return recovery, err
}
