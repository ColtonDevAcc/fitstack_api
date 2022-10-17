package dto

import (
	"time"
)

type UserSignUp struct {
	Id            string    `json:"user_id"`
	Email         string    `json:"email" binding:"required,email"`
	DisplayName   string    `json:"display_name" binding:"required"`
	Password      string    `json:"password"`
	FirstName     string    `json:"first_name" binding:"required"`
	LastName      string    `json:"last_name" binding:"required"`
	PhoneNumber   string    `json:"phone_number" binding:"required"`
	DateOfBirth   string    `json:"date_of_birth" binding:"required"`
	EmailVerified bool      `json:"email_verified"`
	PhoneVerified bool      `json:"phone_verified"`
	PhotoURL      string    `json:"photo_url" binding:"omitempty,url"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}
