package user

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// this is a friendship table struct
type Friendship struct {
	Id           *uuid.UUID `json:"id"`
	FromUser     string     `json:"from_user"`
	ToUser       string     `json:"to_user" binding:"required"`
	Accepted     bool       `json:"accepted"`
	SentTime     time.Time  `json:"sent_time"`
	ResponseTime *time.Time `json:"response_time"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type FriendshipUsecase interface {
	AddFriend(friendship *Friendship) error
	RemoveFriend(ctx context.Context, friendship *Friendship) error
	GetFriends(ctx context.Context, uuid string) ([]*UserProfile, error)
	GetFriendsList(ctx context.Context, uuid string) ([]*UserProfile, error)
	GetFriend(ctx context.Context, email string) (*UserProfile, error)
}

type FriendshipRepository interface {
	AddFriend(friendship *Friendship) error
	GetFriendsList(uuid string) ([]*UserProfile, error)
	RemoveFriend(friendship *Friendship) error
	GetFriends(uuid string) ([]*UserProfile, error)
	GetFriend(email string) (*UserProfile, error)
}
