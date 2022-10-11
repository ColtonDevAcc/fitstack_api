package repository

import (
	"context"
	"fmt"

	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/georgysavva/scany/v2/pgxscan"
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
	VALUES ($1, $2, $3)`
	_, err := f.Database.Exec(context.Background(), insertStatement, &friendship.ToUser, &friendship.FromUser, &friendship.SentTime)
	//TODO: return inserted row
	row := f.Database.QueryRow(context.Background(), insertStatement, &friendship.ToUser, &friendship.FromUser, &friendship.SentTime)

	err = row.Scan(&newFriend)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &newFriend, nil

}

func (f *friendshipRepository) RemoveFriend(friendship *domain.Friendship) error {
	return nil
}

func (f *friendshipRepository) GetFriends(uuid string) ([]*domain.Friendship, error) {
	friendship := []*domain.Friendship{}

	_, err := f.Database.Exec(context.Background(), `DECLARE curs CURSOR WITH HOLD FOR SELECT * FROM friends where to_user = $1;`, uuid)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("user does not exists")
	}
	fetchStatement := `
        FETCH FROM curs;`

	rows, err := f.Database.Query(context.Background(), fetchStatement)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("user does not exists")
	}
	defer rows.Close()
	pgxscan.ScanRow(&friendship, rows)

	f.Database.Exec(context.Background(), `CLOSE curs;`)
	return friendship, nil
}
