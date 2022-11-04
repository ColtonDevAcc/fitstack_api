package repository

import (
	domain "github.com/VooDooStack/FitStackAPI/domain/user"
	"gorm.io/gorm"
)

type friendshipRepository struct {
	Database gorm.DB
}

func NewFriendshipRepository(db gorm.DB) domain.FriendshipRepository {
	return &friendshipRepository{Database: db}
}

func (f *friendshipRepository) AddFriend(friendship *domain.Friendship) error {
	//TODO:
	return nil
}

func (f *friendshipRepository) RemoveFriend(friendship *domain.Friendship) error {
	return nil
}

func (f *friendshipRepository) GetFriends(uuid string) ([]*domain.UserProfile, error) {
	//TODO:
	return nil, nil
}

func (f *friendshipRepository) GetFriendsList(uuid string) ([]*domain.UserProfile, error) {
	//TODO:
	return nil, nil
}

func (f *friendshipRepository) GetFriend(email string) (*domain.UserProfile, error) {
	//TODO:
	return nil, nil
}

func (f *friendshipRepository) getFriends_ToAndFrom(to_user string, from_user string) (*domain.Friendship, error) {
	//TODO:
	return nil, nil
}
