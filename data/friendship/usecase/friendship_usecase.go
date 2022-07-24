package usecase

import "github.com/VooDooStack/FitStackAPI/domain"

type FriendshipUsecase struct {
	friendshipRepo domain.FriendshipRepository
}

func NewFriendshipUsecase(fr domain.FriendshipRepository) domain.FriendshipUsecase {
	return &FriendshipUsecase{friendshipRepo: fr}
}

func (f *FriendshipUsecase) AddFriend(friendship domain.Friendship) (domain.Friendship, error) {
	friendship, err := f.friendshipRepo.AddFriend(friendship)
	if err != nil {
		return domain.Friendship{}, err
	}

	return friendship, nil
}

func (f *FriendshipUsecase) RemoveFriend(friendship domain.Friendship) error {
	err := f.friendshipRepo.RemoveFriend(friendship)
	if err != nil {
		return err
	}

	return nil
}
