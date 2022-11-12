package user

import (
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	ID               string         `json:"id" gorm:"primaryKey; unique"`
	Challenges       []Challenge    `json:"challenges" gorm:"foreignKey:ID"`
	Achievements     []Achievement  `json:"achievements" gorm:"foreignKey:ID"`
	Statistics       UserStatistic  `json:"user_statistics" gorm:"foreignKey:ID"`
	DisplayName      string         `json:"display_name" binding:"required" gorm:"unique"`
	FitCredit        int            `json:"fit_credits"`
	SocialPoints     int            `json:"social_points"`
	DaysLoggedInARow int            `json:"days_logged_in_a_row"`
	Avatar           *string        `json:"avatar"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
