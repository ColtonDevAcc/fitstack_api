package usecase

import (
	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/sirupsen/logrus"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(ur domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: ur}
}

func (u *userUsecase) Delete(uuid string) error {
	err := u.userRepo.Delete(uuid)
	if err != nil {
		logrus.Error(err)

		return err
	}

	return nil
}

func (u *userUsecase) GetByEmail(email string) (domain.User, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		logrus.Error(err)

		return domain.User{DisplayName: "Null User"}, err
	}

	return user, nil
}

func (u *userUsecase) GetByUuid(uuid string) (domain.User, error) {
	user, err := u.userRepo.GetByUuid(uuid)
	if err != nil {
		logrus.Error(err)

		return domain.User{DisplayName: "Null User"}, err
	}

	return user, nil
}

func (u *userUsecase) Store(user domain.User) error {
	err := u.userRepo.Store(user)
	if err != nil {
		logrus.Error(err)

		return err
	}

	return nil
}

func (u *userUsecase) Update(uuid string) error {
	err := u.userRepo.Update(uuid)
	if err != nil {
		logrus.Error(err)

		return err
	}

	return nil
}
