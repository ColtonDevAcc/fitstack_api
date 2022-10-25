package usecase

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"time"

	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"

	"github.com/VooDooStack/FitStackAPI/domain/dto"
	"github.com/VooDooStack/FitStackAPI/domain/user"
	"github.com/sirupsen/logrus"
)

type userUsecase struct {
	userRepo user.UserRepository
	client   auth.Client
	storage  *storage.Client
}

func NewUserUseCase(ur user.UserRepository, client auth.Client, storage *storage.Client) user.UserUsecase {
	return &userUsecase{userRepo: ur, client: client, storage: storage}
}

func (u *userUsecase) Delete(uuid string) error {
	err := u.userRepo.Delete(uuid)
	if err != nil {
		logrus.Error(err)

		return err
	}

	return nil
}

func (u *userUsecase) GetByEmail(email string) (*user.User, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		logrus.Error(err)

		return nil, err
	}

	return user, nil
}

func (u *userUsecase) GetByUuid(uuid string) (*user.User, error) {
	user, err := u.userRepo.GetByUuid(uuid)
	if err != nil {
		logrus.Error(err)

		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Store(user *user.User) error {
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

func (u *userUsecase) SignUp(user *dto.UserSignUp, ctx context.Context) (*user.User, error) {
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

func (u *userUsecase) SignInWithToken(ctx context.Context, token string) (*user.User, error) {
	at, err := u.client.VerifyIDToken(ctx, token)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	user, err := u.userRepo.GetByUuid(at.UID)
	if err != nil {
		logrus.Error(err)
		return nil, err
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

func (u *userUsecase) UpdateUserAvatar(ctx context.Context, uuid string, file *multipart.FileHeader, src io.Reader) (string, error) {
	handler, err := u.storage.Bucket(os.Getenv("BUCKET"))
	if err != nil {
		return "", fmt.Errorf("error getting default bucket")
	}

	sw := handler.Object(file.Filename).NewWriter(context.Background())
	if _, err := io.Copy(sw, src); err != nil {
		return "", fmt.Errorf("error copying bucket: %v", err)
	}

	if err := sw.Close(); err != nil {
		return "", fmt.Errorf("error closing bucket: %v", err)
	}

	avatarUrl, err := url.Parse("/" + os.Getenv("BUCKET") + "/" + sw.Attrs().Name)
	if err != nil {
		return "", fmt.Errorf("error parsing url from bucket %v", err)
	}

	urlString := avatarUrl.String()
	err = u.userRepo.UpdateUserAvatar(uuid, urlString)
	if err != nil {
		logrus.Error(err)
		return "", err
	}

	return urlString, nil
}

func (u *userUsecase) GetUserProfile(uuid string) (*user.UserProfile, error) {
	profile, err := u.userRepo.GetUserProfile(uuid)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return profile, nil
}
