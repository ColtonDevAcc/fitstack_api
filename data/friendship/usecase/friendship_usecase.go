package usecase

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/sirupsen/logrus"
)

type FriendshipUsecase struct {
	friendshipRepo domain.FriendshipRepository
	client         *auth.Client
}

func NewFriendshipUsecase(fr domain.FriendshipRepository, fa *auth.Client) domain.FriendshipUsecase {
	return &FriendshipUsecase{friendshipRepo: fr, client: fa}
}

func (f *FriendshipUsecase) AddFriend(friendship *domain.Friendship) (*domain.Friendship, error) {
	friendship, err := f.friendshipRepo.AddFriend(friendship)
	if err != nil {
		return nil, err
	}

	return friendship, nil
}

func (f *FriendshipUsecase) RemoveFriend(ctx context.Context, friendship *domain.Friendship) error {
	err := f.friendshipRepo.RemoveFriend(friendship)
	if err != nil {
		return err
	}

	return nil
}

func (f *FriendshipUsecase) GetFriends(ctx context.Context, token string) ([]*domain.UserProfile, error) {
	at, err := f.client.VerifyIDToken(ctx, token)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	friendship, err := f.friendshipRepo.GetFriends(at.UID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return friendship, nil
}
