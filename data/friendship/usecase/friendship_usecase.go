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

func (f *FriendshipUsecase) AddFriend(friendship *domain.Friendship) error {
	_, err := f.client.GetUser(context.Background(), friendship.ToUser)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = f.friendshipRepo.AddFriend(friendship)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (f *FriendshipUsecase) RemoveFriend(ctx context.Context, friendship *domain.Friendship) error {
	err := f.friendshipRepo.RemoveFriend(friendship)
	if err != nil {
		return err
	}

	return nil
}

func (f *FriendshipUsecase) GetFriends(ctx context.Context, uuid string) ([]*domain.UserProfile, error) {
	friendship, err := f.friendshipRepo.GetFriends(uuid)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return friendship, nil
}

func (f *FriendshipUsecase) GetFriendsList(ctx context.Context, uuid string) ([]*domain.UserProfile, error) {
	friendship, err := f.friendshipRepo.GetFriendsList(uuid)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return friendship, nil
}

func (f *FriendshipUsecase) GetFriend(ctx context.Context, email string) (*domain.UserProfile, error) {
	friendship, err := f.friendshipRepo.GetFriend(email)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return friendship, nil
}
