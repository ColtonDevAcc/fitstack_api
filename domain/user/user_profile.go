package user

import (
	"time"
)

type UserProfile struct {
	Id               string           `json:"id"`
	Challenges       []*Challenge     `json:"challenges"`
	Achievements     []*Achievement   `json:"achievements"`
	Statistics       []*UserStatistic `json:"user_statistics"`
	Friends          []*UserProfile   `json:"friends"`
	DisplayName      string           `json:"display_name" binding:"required"`
	FitCredit        int              `json:"fit_credits"`
	SocialPoints     int              `json:"social_points"`
	DaysLoggedInARow int              `json:"days_logged_in_a_row"`
	UpdatedAt        *time.Time       `json:"updated_at"`
	Avatar           *string          `json:"avatar"`
	Accepted         *bool            `json:"accepted"`
}
