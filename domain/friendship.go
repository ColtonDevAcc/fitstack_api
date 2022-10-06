package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// this is a friendship table struct
type Friendship struct {
	gorm.Model
	FromUserId   string    `gorm:"primaryKey" json:"from_user"`
	ToUserId     string    `gorm:"primaryKey" json:"to_user" binding:"required"`
	SentTime     time.Time `json:"sent_time"`
	ResponseTime time.Time `json:"response_time"`
	Accepted     bool      `json:"accepted"`
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
