package usecase

import (
	"context"
	"time"

	"firebase.google.com/go/v4/auth"

	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/VooDooStack/FitStackAPI/domain/dto"
	"github.com/sirupsen/logrus"
)

type userUsecase struct {
	userRepo domain.UserRepository
	client   auth.Client
}

func NewUserUseCase(ur domain.UserRepository, client auth.Client) domain.UserUsecase {
	return &userUsecase{userRepo: ur, client: client}
}

func (u *userUsecase) Delete(uuid string) error {
	err := u.userRepo.Delete(uuid)
	if err != nil {
		logrus.Error(err)

		return err
	}

	return nil
}

func (u *userUsecase) GetByEmail(email string) (*domain.User, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		logrus.Error(err)

		return &domain.User{DisplayName: "Null User"}, err
	}

	return user, nil
}

func (u *userUsecase) GetByUuid(uuid string) (*domain.User, error) {
	user, err := u.userRepo.GetByUuid(uuid)
	if err != nil {
		logrus.Error(err)

		return &domain.User{DisplayName: "Null User"}, err
	}

	return user, nil
}

func (u *userUsecase) Store(user *domain.User) error {
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

func (u *userUsecase) SignUp(user *dto.UserSignUp, ctx context.Context) (*domain.User, error) {
	params := (&auth.UserToCreate{}).Email(user.Email).Password(user.Password).PhoneNumber(user.PhoneNumber).DisplayName(user.DisplayName)

	err := u.userRepo.CheckUniqueFields(user)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	fbu, err := u.client.CreateUser(ctx, params)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	user.Id = fbu.UID
	user.CreatedAt = time.Now()

	newUser, err := u.userRepo.SignUp(user)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return newUser, nil
}

func (u *userUsecase) SignInWithToken(ctx context.Context, token string) (*domain.User, error) {
	at, err := u.client.VerifyIDToken(ctx, token)
	if err != nil {
		logrus.Error(err)
		return &domain.User{}, err
	}

	user, err := u.userRepo.GetByUuid(at.UID)
	if err != nil {
		logrus.Error(err)
		return &domain.User{}, err
	}

	return user, nil
}

func (u *userUsecase) RefreshToken(ctx context.Context, refresh_token string) (string, error) {
	str, err := u.userRepo.RefreshToken(refresh_token)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return str, err
}

func (u *userUsecase) SignInWithEmailAndPassword(ctx context.Context, login *dto.LoginInEmailAndPassword) (string, error) {
	response, err := u.userRepo.SignInWithEmailAndPassword(login)
	if err != nil {
		logrus.Error(err)
		return err.Error(), err
	}

	return response, nil
}
