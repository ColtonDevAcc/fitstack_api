package repository

import (
	"github.com/VooDooStack/FitStackAPI/domain/dto"
	"github.com/VooDooStack/FitStackAPI/domain/user"
	"gorm.io/gorm"
)

type userRepository struct {
	Database         gorm.DB
	FriendRepository user.FriendshipRepository
}

func NewUserRepository(db gorm.DB, fr user.FriendshipRepository) user.UserRepository {
	return &userRepository{db, fr}
}

func (u *userRepository) Delete(uuid string) error {
	//TODO:

	return nil
}

func (u *userRepository) GetByEmail(email string) (*user.User, error) {
	//TODO:

	return nil, nil
}

func (u *userRepository) GetByUuid(uuid string) (*user.User, error) {
	ur := &user.User{}
	err := u.Database.Where("id = ?", uuid).Preload("Profile").Preload("Friends").First(&ur).Error
	return ur, err
}

func (u *userRepository) Store(user *user.User) error {
	//TODO:

	return nil
}

func (u *userRepository) Update(uuid string) error {
	//TODO:
	return nil
}

func (u *userRepository) SignUp(user *dto.UserSignUp) (*user.User, error) {
	//TODO:

	return nil, nil
}

func (u *userRepository) RefreshToken(refresh_token string) (string, error) {
	//TODO:
	return "sb", nil
}

func (u *userRepository) SignInWithEmailAndPassword(login *dto.LoginInEmailAndPassword) (string, error) {
	//TODO:
	return "", nil
}

func (u *userRepository) CheckUniqueFields(user *dto.UserSignUp) error {
	//TODO:

	return nil
}

func (u *userRepository) UpdateUserAvatar(uuid string, fileURL string) error {
	//TODO:
	return nil
}

func (u *userRepository) GetUserProfile(uuid string) (*user.UserProfile, error) {
	//TODO:

	return nil, nil
}
