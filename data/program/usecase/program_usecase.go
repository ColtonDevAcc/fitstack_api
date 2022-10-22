package usecase

import (
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"

	"github.com/VooDooStack/FitStackAPI/domain"
)

type programUsecase struct {
	userRepo domain.UserRepository
	client   auth.Client
	storage  *storage.Client
}
