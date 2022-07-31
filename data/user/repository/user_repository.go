package repository

import (
	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepository struct {
	Database gorm.DB
}

func NewUserRepository(db gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (u *userRepository) Delete(uuid string) error {
	tx := u.Database.Where(domain.User{UserId: uuid}).Delete(&domain.User{})
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return tx.Error
	}

	return nil
}

func (u *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	tx := u.Database.Where(domain.User{Email: email}).Find(&user)
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return domain.User{DisplayName: "Null User"}, tx.Error
	}

	return user, nil
}

func (u *userRepository) GetByUuid(uuid string) (domain.User, error) {
	var user domain.User
	tx := u.Database.Where(domain.User{UserId: uuid}).Find(&user)
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return domain.User{DisplayName: "Null User"}, tx.Error
	}

	return user, nil
}

func (u *userRepository) Store(user domain.User) error {
	tx := u.Database.Create(&user)
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return tx.Error
	}

	return nil
}

func (u *userRepository) Update(uuid string) error {
	tx := u.Database.Where(domain.User{UserId: uuid}).Save(&domain.User{})
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return tx.Error
	}

	return nil
}

func (u *userRepository) SignUp(user domain.User) (domain.User, error) {
	tx := u.Database.Create(&user)
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return domain.User{DisplayName: "Null User"}, tx.Error
	}

	return user, nil
}
