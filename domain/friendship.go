package domain

import (
	"context"
	"time"
)

// this is a friendship table struct
type Friendship struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	FromUserId   string    `json:"from_user"`
	ToUserId     string    `json:"to_user" binding:"required"`
	SentTime     time.Time `json:"sent_time"`
	ResponseTime time.Time `json:"response_time"`
	Accepted     bool      `json:"accepted" `
	User         User      `gorm:"many2many:friends"`
}

type FriendshipUsecase interface {
	AddFriend(friendship Friendship) (Friendship, error)
	RemoveFriend(ctx context.Context, friendship Friendship) error
	GetFriends(ctx context.Context, token string) ([]Friendship, error)
}

type FriendshipRepository interface {
	AddFriend(friendship Friendship) (Friendship, error)
	RemoveFriend(friendship Friendship) error
	GetFriends(uuid string) ([]Friendship, error)
}
