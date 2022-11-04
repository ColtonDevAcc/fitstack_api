package user

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// this is a friendship table struct
type Friendship struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	FromUserID   uint           `json:"from_user_id"`
	FromUser     User           `json:"from_user" binding:"required" gorm:"foreignKey:FromUserID"`
	ToUserID     uint           `json:"to_user_id"`
	ToUser       User           `json:"to_user" binding:"required" gorm:"foreignKey:ToUserID"`
	Accepted     bool           `json:"accepted"`
	SentTime     time.Time      `json:"sent_time"`
	ResponseTime *time.Time     `json:"response_time"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
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
