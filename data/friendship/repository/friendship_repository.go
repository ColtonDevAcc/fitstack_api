package repository

import (
	"github.com/VooDooStack/FitStackAPI/domain"
	"gorm.io/gorm"
)

type friendshipRepository struct {
	Database gorm.DB
}

func NewFriendshipRepository(db gorm.DB) domain.FriendshipRepository {
	return &friendshipRepository{Database: db}
}

func (f *friendshipRepository) AddFriend(friendship domain.Friendship) (domain.Friendship, error) {
	tx := f.Database.Create(&friendship)
	if tx.Error != nil {
		return domain.Friendship{}, tx.Error
	}

	return friendship, nil
}

func (f *friendshipRepository) RemoveFriend(friendship domain.Friendship) error {
	return nil
}
