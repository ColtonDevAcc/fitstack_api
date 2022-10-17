package domain

import "time"

type UserProfile struct {
	Id               string           `json:"id"`
	Challenges       []*Challenge     `json:"challenges"`
	Achievements     []*Achievement   `json:"achievements"`
	Statistics       []*UserStatistic `json:"user_statistics"`
	FitCredit        int              `json:"fit_credits"`
	SocialPoints     int              `json:"social_points"`
	DaysLoggedInARow int              `json:"days_logged_in_a_row"`
	UpdatedAt        *time.Time       `json:"updated_at"`
}
