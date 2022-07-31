package domain

// this is a friendship table struct
type Friendship struct {
	ID                      int    `gorm:"primaryKey"`
	Accepted                bool   `json:"accepted"`
	FriendshipRequestTarget string `gorm:"references:UserUuid"`
	FriendshipRequestedBy   string `gorm:"references:UserUuid"`
}

type FriendshipUsecase interface {
	AddFriend(friendship Friendship) (Friendship, error)
	RemoveFriend(friendship Friendship) error
}

type FriendshipRepository interface {
	AddFriend(friendship Friendship) (Friendship, error)
	RemoveFriend(friendship Friendship) error
}
