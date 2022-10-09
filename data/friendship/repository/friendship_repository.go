package repository

import (
	"context"
	"fmt"

	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type friendshipRepository struct {
	Database pgx.Conn
}

func NewFriendshipRepository(db pgx.Conn) domain.FriendshipRepository {
	return &friendshipRepository{Database: db}
}

func (f *friendshipRepository) AddFriend(friendship *domain.Friendship) (*domain.Friendship, error) {
	var id *string
	userCheckSql := `
	SELECT * FROM users
	WHERE id='$1';
	RETURNING id
	`
	err := f.Database.QueryRow(context.Background(), userCheckSql, &friendship.ToUserId).Scan(&id)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("user does not exists")
	}

	insertStatement := `
	INSERT INTO friends ( to_user, from_user, updated_at, accepted, sent_time, deleted_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING *
	`
	rows, _ := f.Database.Query(context.Background(), insertStatement, &friendship.ToUserId, &friendship.FromUserId, &friendship.UpdatedAt, &friendship.Accepted, &friendship.SentTime, &friendship.DeletedAt)
	defer rows.Close()
	pgxscan.ScanRow(&friendship, rows)

	return friendship, nil

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
