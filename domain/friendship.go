package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// this is a friendship table struct
type Friendship struct {
	gorm.Model
	ID           int       `gorm:"primaryKey" json:"id"`
	FromUserId   string    `json:"from_user"`
	ToUserId     string    `json:"to_user" binding:"required"`
	SentTime     time.Time `json:"sent_time"`
	ResponseTime time.Time `json:"response_time"`
	Accepted     bool      `json:"accepted" `
	User         *User     `gorm:"foreignKey:FromUserId;references:UserId;foreignKey:ToUserId;references:UserId" json:"-"`
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
