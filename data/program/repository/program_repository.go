package repository

import (
	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type programRepository struct {
	Database         pgxpool.Pool
	FriendRepository domain.FriendshipRepository
}

func NewProgramRepository(db pgxpool.Pool, fr domain.FriendshipRepository) domain.ProgramRepository {
	return &programRepository{db, fr}
}

func (u *programRepository) Delete(uuid string) error {

	return nil
}
