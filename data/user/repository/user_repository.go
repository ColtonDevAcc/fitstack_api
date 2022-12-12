package repository

import (
	"time"

	"github.com/VooDooStack/FitStackAPI/domain/dto"
	healthLogs "github.com/VooDooStack/FitStackAPI/domain/health_logs"
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
	err := u.Database.Where("id = ?", uuid).Preload("Profile.Statistics.WeightLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at desc").Limit(1)
	}).Preload("Profile.Statistics.BodyFatPercentageLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at desc").Limit(1)
	}).Preload("Profile.Statistics.BodyMassIndexLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at desc").Limit(1)
	}).Preload("Profile.Statistics.ActiveEnergyBurnedLogs", func(tx *gorm.DB) *gorm.DB {
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
	err := u.Database.Model(&user.UserProfile{}).Where("id = ?", uuid).Update("avatar", fileURL).Error
	return err
}

func (u *userRepository) GetUserProfile(uuid string) (*user.UserProfile, error) {
	ur := &user.UserProfile{}
	err := u.Database.Where("id = ?", uuid).First(&ur).Error
	return ur, err
}

func (u *userRepository) UpdateUserProfile(uuid string, profile *user.UserProfile) error {
	return u.Database.Model(&user.UserProfile{}).Where("id = ?", uuid).Updates(profile).Error
}

func (u *userRepository) UpdateUserStatistics(userStatistic *user.UserStatistic) error {
	err := u.Database.Model(&user.UserStatistic{}).Where("id = ?", userStatistic.ID).Updates(userStatistic).Error
	if err != nil {
		return err
	}

	if userStatistic.WeightLogs != nil && len(userStatistic.WeightLogs) > 0 {
		err = u.Database.CreateInBatches(&userStatistic.WeightLogs, 1000).Error
	}

	if userStatistic.BodyMassIndexLogs != nil && len(userStatistic.BodyMassIndexLogs) > 0 {
		err = u.Database.CreateInBatches(&userStatistic.BodyMassIndexLogs, 1000).Error
	}

	if userStatistic.BodyFatPercentageLogs != nil && len(userStatistic.BodyFatPercentageLogs) > 0 {
		err = u.Database.CreateInBatches(&userStatistic.BodyFatPercentageLogs, 1000).Error
	}

	if userStatistic.ActiveEnergyBurnedLogs != nil && len(userStatistic.ActiveEnergyBurnedLogs) > 0 {
		err = u.Database.CreateInBatches(&userStatistic.ActiveEnergyBurnedLogs, 1000).Error
	}

	if userStatistic.StepsLogs != nil && len(userStatistic.StepsLogs) > 0 {
		err = u.Database.CreateInBatches(&userStatistic.StepsLogs, 1000).Error
	}

	if userStatistic.SleepAwakeLogs != nil && len(userStatistic.SleepAwakeLogs) > 0 {
		err = u.Database.CreateInBatches(&userStatistic.SleepAwakeLogs, 1000).Error
	}

	return err
}

func (u *userRepository) GetUserStatistics(uuid string) (*user.UserStatistic, error) {
	us := user.UserStatistic{}
	err := u.Database.Model(user.UserStatistic{}).Where("id = ?", uuid).Preload("HeightLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at desc").Limit(1)
	}).Preload("WeightLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("BodyMassIndexLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("BodyFatPercentageLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("StepsLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("HeartRateLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("SleepLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("BloodPressureLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("BodyTemperatureLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("RespiratoryRateLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("OxygenSaturationLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("ActiveEnergyBurnedLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("BasalEnergyBurnedLog", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).First(&us).Error

	return &us, err
}

func (u *userRepository) GetUserStatisticsSnapshot(uuid string) (*user.UserStatistic, error) {
	us := user.UserStatistic{}
	err := u.Database.Where("id = ?", uuid).Preload("WeightLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("BodyMassIndexLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("BodyFatPercentageLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("StepsLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("SleepInBedLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("SleepAwakeLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("ActiveEnergyBurnedLogs", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).Preload("BasalEnergyBurnedLog", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("DISTINCT ON (date_trunc('day', created_at)) *").Where("created_at > ?", time.Now().AddDate(0, -1, 0)).Order("date_trunc('day', created_at), created_at asc").Limit(30)
	}).First(&us).Error
	return &us, err
}

func (u *userRepository) GetUserHealthLog(uuid string, healthType healthLogs.HealthDataType) (*user.UserStatistic, error) {
	us := user.UserStatistic{}
	switch healthType {
	case healthLogs.WEIGHT:
		err := u.Database.Model(&user.UserStatistic{}).Where("id = ?", uuid).Preload("WeightLogs").Statement.Order("created_at desc").Take(&us).Error
		return &us, err
	case healthLogs.BODY_FAT_PERCENTAGE:
		err := u.Database.Model(&user.UserStatistic{}).Where("id = ?", uuid).Preload("BodyFatPercentageLogs").Statement.Order("created_at desc").Take(&us).Error
		return &us, err
	case healthLogs.BODY_MASS_INDEX:
		err := u.Database.Model(&user.UserStatistic{}).Where("id = ?", uuid).Preload("BodyMassIndexLogs").Statement.Order("created_at desc").Take(&us).Error
		return &us, err
	case healthLogs.STEPS:
		err := u.Database.Model(&user.UserStatistic{}).Where("id = ?", uuid).Preload("StepsLogs").Statement.Order("created_at desc").Take(&us).Error
		return &us, err
	case healthLogs.DISTANCE_WALKING_RUNNING:
		err := u.Database.Model(&user.UserStatistic{}).Where("id = ?", uuid).Preload("DistanceLogs").Statement.Order("created_at desc").Take(&us).Error
		return &us, err
	case healthLogs.MOVE_MINUTES:
		err := u.Database.Model(&user.UserStatistic{}).Where("id = ?", uuid).Preload("ActiveMinutesLogs").Statement.Order("created_at desc").Take(&us).Error
		return &us, err
	case healthLogs.SLEEP_IN_BED:
		err := u.Database.Model(&user.UserStatistic{}).Where("id = ?", uuid).Preload("SleepLogs").Statement.Order("created_at desc").Take(&us).Error
		return &us, err
	case healthLogs.ACTIVE_ENERGY_BURNED:
		err := u.Database.Model(&user.UserStatistic{}).Where("id = ?", uuid).Preload("ActiveEnergyBurnedLogs").Statement.Order("created_at desc").Take(&us).Error
		return &us, err
	case healthLogs.BASAL_ENERGY_BURNED:
		err := u.Database.Model(&user.UserStatistic{}).Where("id = ?", uuid).Preload("BasalEnergyBurnedLog").Statement.Order("created_at desc").Take(&us).Error
		return &us, err
	default: // healthType.ALL
		err := u.Database.Model(&user.UserStatistic{}).Where("id = ?", uuid).Preload("WeightLogs").Preload("BodyMassIndexLogs").Preload("BodyFatPercentageLogs").Preload("StepsLogs").Preload("DistanceLogs").Preload("ActiveMinutesLogs").Preload("SleepLogs").Preload("ActiveEnergyBurnedLogs").Preload("BasalEnergyBurnedLog").Statement.Order("created_at desc").Take(&us).Error
		return &us, err
	}
}
