package repository

import (
	"context"

	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type friendshipRepository struct {
	Database pgxpool.Pool
}

func NewFriendshipRepository(db pgxpool.Pool) domain.FriendshipRepository {
	return &friendshipRepository{Database: db}
}

func (f *friendshipRepository) AddFriend(friendship *domain.Friendship) (*domain.Friendship, error) {
	newFriend := domain.Friendship{}
	insertStatement := `
	INSERT INTO friends (to_user, from_user, sent_time)
	VALUES ($1, $2, $3)
	RETURNING *
	`
	row := f.Database.QueryRow(context.Background(), insertStatement, &friendship.ToUser, &friendship.FromUser, &friendship.SentTime)
	err := row.Scan(&newFriend)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &newFriend, nil

}

func (f *friendshipRepository) RemoveFriend(friendship *domain.Friendship) error {
	// tx := f.Database.Where(friendship).First(&friendship).Delete(&friendship)
	// if tx.Error != nil {
	// 	logrus.Error(tx.Error)

	// 	return tx.Error
	// }

	return nil

}

func (f *friendshipRepository) GetFriends(uuid string) ([]*domain.Friendship, error) {
	// var friendshipList []domain.Friendship
	// tx := f.Database.Model(domain.Friendship{}).Where(domain.Friendship{ToUserId: uuid}).Or(domain.Friendship{FromUserId: uuid}).Find(&friendshipList)
	// if tx.Error != nil {
	// 	logrus.Error(tx.Error)
	// 	return []domain.Friendship{}, tx.Error
	// }
	// fmt.Printf("============================================\nfriendships: %v \n============================================\n", friendshipList)
	return nil, nil
}
