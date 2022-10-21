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

func (f *friendshipRepository) AddFriend(friendship *domain.Friendship) error {
	friend, err := f.getFriends_ToAndFrom(friendship.ToUser, friendship.FromUser)
	if err != nil {
		logrus.Error(err)
		return err
	}

	if friend != (&domain.Friendship{ToUser: friendship.ToUser, FromUser: friend.FromUser}) {
		insertStatement := `
		INSERT INTO friends (to_user, from_user)
		VALUES 
		((SELECT id from users WHERE id= $1 ), 
		(SELECT id from users WHERE id= $2 ))`
		_, err = f.Database.Exec(context.Background(), insertStatement, &friendship.ToUser, &friendship.FromUser)
		if err != nil {
			logrus.Error(err)
			return err
		}

	} else {
		logrus.Print("friendship already exists")
		return fmt.Errorf("friendship already exists")
	}

	return nil
}

func (f *friendshipRepository) RemoveFriend(friendship *domain.Friendship) error {
	return nil
}

func (f *friendshipRepository) GetFriends(uuid string) ([]*domain.UserProfile, error) {
	friendship := []*domain.UserProfile{}
	queryStatement :=
		`
	SELECT DISTINCT
	u.id, u.challenges, u.achievements, u.statistics, u.fit_credit, u.social_points, u.days_logged_in_a_row, u.display_name, u.updated_at, u.avatar
	FROM
	friends f
	JOIN user_profiles u
    on f.to_user = u.id AND f.accepted = true OR f.from_user = u.id AND f.accepted = true
	WHERE f.from_user = $1 AND U.id != $1 OR f.to_user= $1 AND U.id != $1
	`

	rows, err := f.Database.Query(context.Background(), queryStatement, uuid)
	if err != nil {
		rows.Close()
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()
	err = pgxscan.ScanAll(&friendship, rows)
	if err != nil {
		rows.Close()
		logrus.Error(err)
		return nil, err
	}

	return friendship, nil
}

func (f *friendshipRepository) GetFriendsList(uuid string) ([]*domain.UserProfile, error) {
	friendship := []*domain.UserProfile{}
	queryStatement :=
		`
	SELECT DISTINCT
	u.id, u.challenges, u.achievements, u.statistics, u.fit_credit, u.social_points, u.days_logged_in_a_row, u.display_name, u.updated_at, u.avatar, f.accepted
	FROM
	friends f
	JOIN user_profiles u
    on f.to_user = u.id OR f.from_user = u.id
	WHERE f.from_user =$1 AND U.id != $1 OR f.to_user=$1 AND U.id != $1
	`

	rows, err := f.Database.Query(context.Background(), queryStatement, uuid)
	if err != nil {
		rows.Close()
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()
	err = pgxscan.ScanAll(&friendship, rows)
	if err != nil {
		rows.Close()
		logrus.Error(err)
		return nil, err
	}

	return friendship, nil
}

func (f *friendshipRepository) GetFriend(email string) (*domain.UserProfile, error) {
	friendship := domain.UserProfile{}
	sqlStatement := `
	SELECT up.id, up.avatar, up.display_name, up.achievements, up.challenges, up.days_logged_in_a_row, up.fit_credit, up.social_points, up.statistics, up.updated_at FROM users
    JOIN user_profiles up
    on up.id = users.id
	WHERE email=$1;
	`
	rows, err := f.Database.Query(context.Background(), sqlStatement, email)
	if err != nil {
		logrus.Error("error querying database error: %v", err)
		return nil, err
	}

	defer rows.Close()
	err = pgxscan.ScanOne(&friendship, rows)
	if err != nil {
		logrus.Error("error scanning row error: %v", err)
		return nil, err
	}

	return &friendship, nil
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
