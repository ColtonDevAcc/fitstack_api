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
	var newFriend *domain.Friendship
	friend, err := f.getFriends_ToAndFrom(friendship.ToUser, friendship.FromUser)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	if *friend == (domain.Friendship{}) {
		insertStatement := `
		INSERT INTO friends (to_user, from_user)
		VALUES ((SELECT id from users WHERE id= $1 ), (SELECT id from users WHERE id= $2 )) ON CONFLICT DO NOTHING`
		_, err = f.Database.Exec(context.Background(), insertStatement, friendship.ToUser, friendship.FromUser)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		newFriend, err = f.getFriends_ToAndFrom(friendship.ToUser, friendship.FromUser)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

	} else {
		fmt.Printf("here is your user %v", friend)
		return nil, fmt.Errorf("friendship already exists")
	}

	return newFriend, nil
}

func (f *friendshipRepository) RemoveFriend(friendship *domain.Friendship) error {
	return nil
}

func (f *friendshipRepository) GetFriends(uuid string) ([]*domain.User, error) {
	friendship := []*domain.User{}
	queryStatement := `WITH vars (toId, fromId) AS (
  	VALUES ((SELECT to_user FROM friends WHERE to_user = $1 AND accepted =true),
    (SELECT from_user FROM friends WHERE from_user = $1 AND accepted =true))
	)
	SELECT id,display_name, first_name, last_name, date_of_birth, photo_url, created_at FROM users, vars WHERE id = vars.toId OR id = vars.fromId
	`

	rows, err := f.Database.Query(context.Background(), queryStatement, uuid)
	if err != nil {
		rows.Close()
		logrus.Error(err)
		return nil, err
	}

	pgxscan.ScanAll(&friendship, rows)

	return friendship, nil
}

func (f *friendshipRepository) getFriends_ToAndFrom(to_user string, from_user string) (*domain.Friendship, error) {
	friend := domain.Friendship{}
	rows, err := f.Database.Query(context.Background(), `SELECT * FROM friends WHERE to_user = $1 AND from_user = $2`, to_user, from_user)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()
	pgxscan.ScanOne(&friend, rows)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &friend, nil
}
