package domain

import (
	"context"
	"io"
	"mime/multipart"
	"time"

	"github.com/VooDooStack/FitStackAPI/domain/dto"
)

type User struct {
	Id            string     `json:"user_id"`
	Email         string     `json:"email" binding:"required,email"`
	DisplayName   string     `json:"display_name" binding:"required"`
	FirstName     string     `json:"first_name" binding:"required"`
	LastName      string     `json:"last_name" binding:"required"`
	PhoneNumber   string     `json:"phone_number" binding:"required"`
	DateOfBirth   string     `json:"date_of_birth" binding:"required"`
	EmailVerified bool       `json:"email_verified"`
	PhoneVerified bool       `json:"phone_verified"`
	PhotoURL      *string    `json:"photo_url" binding:"omitempty,url"`
	UpdatedAt     *time.Time `json:"updated_at"`
	CreatedAt     *time.Time `json:"created_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
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
}
