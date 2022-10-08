package repository

import (
	"fmt"

	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type friendshipRepository struct {
	Database pgx.Conn
}

func NewFriendshipRepository(db pgx.Conn) domain.FriendshipRepository {
	return &friendshipRepository{Database: db}
}

func (f *friendshipRepository) AddFriend(friendship domain.Friendship) (domain.Friendship, error) {
	var toFriend domain.User
	tx := f.Database.Where(domain.User{UUID: friendship.ToUserId}).First(&toFriend)
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return domain.Friendship{}, fmt.Errorf("this user does not exist: %s", friendship.ToUserId)
	}

	var friend domain.Friendship
	tx = f.Database.Where(domain.Friendship{ToUserId: friend.ToUserId}).First(&friend)
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return domain.Friendship{}, tx.Error
	}

	// if friend == (domain.Friendship{}) {
	// 	tx := f.Database.Create(&friendship)
	// 	if tx.Error != nil {
	// 		return domain.Friendship{}, tx.Error
	// 	}

	// 	return friendship, nil
	// }

	return domain.Friendship{}, fmt.Errorf("an invitation to %q already exists", friendship.ToUserId)
}

func (f *friendshipRepository) RemoveFriend(friendship domain.Friendship) error {
	tx := f.Database.Where(friendship).First(&friendship).Delete(&friendship)
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return tx.Error
	}

	return nil

}

func (f *friendshipRepository) GetFriends(uuid string) ([]domain.Friendship, error) {
	var friendshipList []domain.Friendship
	tx := f.Database.Model(domain.Friendship{}).Where(domain.Friendship{ToUserId: uuid}).Or(domain.Friendship{FromUserId: uuid}).Find(&friendshipList)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return []domain.Friendship{}, tx.Error
	}
	fmt.Printf("============================================\nfriendships: %v \n============================================\n", friendshipList)
	return friendshipList, nil
}
