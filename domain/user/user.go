package user

import (
	"context"
	"io"
	"mime/multipart"
	"time"

	"github.com/VooDooStack/FitStackAPI/domain/dto"
	"gorm.io/gorm"
)

type User struct {
	ID            string         `json:"id" binding:"required" gorm:"primaryKey;unique;not null;"`
	Email         string         `json:"email" binding:"required,email" gorm:"unique;not null;"`
	FirstName     string         `json:"first_name" binding:"required" gorm:"not null;"`
	LastName      string         `json:"last_name" binding:"required" gorm:"not null;"`
	PhoneNumber   string         `json:"phone_number" binding:"required" gorm:"unique,not null"`
	DateOfBirth   time.Time      `json:"date_of_birth" binding:"required"`
	EmailVerified bool           `json:"email_verified"`
	PhoneVerified bool           `json:"phone_verified"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime:true"`
	CreatedAt     time.Time      `json:"created_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	WeightGoal    float64        `json:"weight_goal"`
	BMIGoal       float64        `json:"bmi_goal"`
	Profile       UserProfile    `json:"profile" gorm:"foreignKey:ID;References:ID"`
	Friends       []Friendship   `json:"friends" gorm:"many2many:user_friends;"`
}

type UserUsecase interface {
	SignUp(user *dto.UserSignUp, ctx context.Context) (*User, error)
	SignInWithToken(ctx context.Context, token string) (*User, error)
	UpdateUserAvatar(ctx context.Context, uuid string, file *multipart.FileHeader, src io.Reader) (string, error)
	SignInWithEmailAndPassword(ctx context.Context, login *dto.LoginInEmailAndPassword) (string, error)
	RefreshToken(ctx context.Context, refresh_token string) (string, error)
	GetByUuid(uuid string) (*User, error)
	Update(uuid string) error
	GetByEmail(email string) (*User, error)
	Store(user *User) error
	Delete(uuid string) error
	GetUserProfile(uuid string) (*UserProfile, error)
	UpdateUserStatistics(userStatistic *UserStatistic) error
	GetUserStatistics(uuid string) (*UserStatistic, error)
	GetUserStatisticsSnapshot(uuid string) (*UserStatistic, error)
}

type UserRepository interface {
	SignUp(user *dto.UserSignUp) (*User, error)
	SignInWithEmailAndPassword(login *dto.LoginInEmailAndPassword) (string, error)
	UpdateUserAvatar(uuid string, fileURL string) error
	RefreshToken(refresh_token string) (string, error)
	GetByUuid(uuid string) (*User, error)
	Update(uuid string) error
	GetByEmail(email string) (*User, error)
	Store(user *User) error
	Delete(uuid string) error
	CheckUniqueFields(user *dto.UserSignUp) error
	GetUserProfile(uuid string) (*UserProfile, error)
	UpdateUserStatistics(userStatistic *UserStatistic) error
	GetUserStatistics(uuid string) (*UserStatistic, error)
	GetUserStatisticsSnapshot(uuid string) (*UserStatistic, error)
}
