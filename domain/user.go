package domain

import (
	"context"
	"time"

	"github.com/VooDooStack/FitStackAPI/domain/dto"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id            uuid.UUID      `gorm:"primaryKey;unique;not null" json:"user_id"`
	Email         string         `gorm:"unique;not null" json:"email" binding:"required,email"`
	Password      string         `gorm:"-:all" json:"password"`
	DisplayName   string         `gorm:"unique;not null" json:"display_name" binding:"required"`
	FirstName     string         `gorm:"not null" json:"first_name" binding:"required"`
	LastName      string         `gorm:"not null" json:"last_name" binding:"required"`
	PhoneNumber   string         `gorm:"unique;not null" json:"phone_number" binding:"required"`
	DateOfBirth   string         `gorm:"not null" json:"date_of_birth" binding:"required"`
	EmailVerified bool           `json:"email_verified"`
	PhoneVerified bool           `json:"phone_verified"`
	PhotoURL      string         `json:"photo_url" binding:"omitempty,url"`
	UpdatedAt     time.Time      `json:"updated_at"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	RefreshToken  string         `json:"refresh_token"`
	Friendship    []Friendship   `gorm:"many2many:users_friendships;"`
}

type UserUsecase interface {
	SignUp(user *User, ctx context.Context) (*User, error)
	SignInWithToken(ctx context.Context, token string) (*User, error)
	SignInWithEmailAndPassword(ctx context.Context, login *dto.LoginInEmailAndPassword) (string, error)
	RefreshToken(ctx context.Context, refresh_token string) (string, error)
	GetByUuid(uuid string) (*User, error)
	Update(uuid string) error
	GetByEmail(email string) (*User, error)
	Store(user *User) error
	Delete(uuid string) error
}

type UserRepository interface {
	SignUp(user *User) (*User, error)
	SignInWithToken(uuid string) (*User, error)
	SignInWithEmailAndPassword(login *dto.LoginInEmailAndPassword) (string, error)
	RefreshToken(refresh_token string) (string, error)
	GetByUuid(uuid string) (*User, error)
	Update(uuid string) error
	GetByEmail(email string) (*User, error)
	Store(user *User) error
	Delete(uuid string) error
	CheckUniqueFields(user *User) error
}
