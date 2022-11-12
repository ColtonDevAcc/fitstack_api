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
	err := u.Database.Where("id = ?", uuid).Preload("Profile.Statistics.HeightLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at desc").Limit(1)
	}).Preload("Profile.Statistics.BodyFatLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at desc").Limit(1)
	}).Preload("Profile.Statistics.BMILogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at desc").Limit(1)
	}).Preload("Profile.Statistics.WeightLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at desc").Limit(1)
	}).Preload("Friends").First(&ur).Error
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

func (u *userRepository) UpdateUserProfile(uuid string, profile *user.UserProfile) error {
	return u.Database.Model(&user.UserProfile{}).Where("id = ?", uuid).Updates(profile).Error
}

func (u *userRepository) UpdateUserStatistics(userStatistic *user.UserStatistic) error {
	err := u.Database.Model(&user.HeightLog{}).Where("user_statistic_id = ?", &userStatistic.ID).Create(&userStatistic.HeightLogs).Error
	if err != nil {
		return err
	}

	err = u.Database.Model(&user.WeightLog{}).Where("user_statistic_id = ?", &userStatistic.ID).Create(&userStatistic.WeightLogs).Error
	if err != nil {
		return err
	}

	err = u.Database.Model(&user.BMILog{}).Where("user_statistic_id = ?", &userStatistic.ID).Create(&userStatistic.BMILogs).Error
	if err != nil {
		return err
	}

	err = u.Database.Model(&user.BodyFatLog{}).Where("user_statistic_id = ?", &userStatistic.ID).Create(&userStatistic.BodyFatLogs).Error
	if err != nil {
		return err
	}

	return err
}

func (u *userRepository) GetUserStatistics(uuid string) (*user.UserStatistic, error) {
	us := user.UserStatistic{}
	err := u.Database.Where("id = ?", uuid).Preload("HeightLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at desc")
	}).Preload("WeightLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at desc")
	}).Preload("BMILogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at desc")
	}).Preload("BodyFatLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at desc")
	}).First(&us).Error

	return &us, err
}
