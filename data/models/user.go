package models

type User struct {
	ID          int    `json:"id"`
	Uuid        string `json:"uuid"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   int64  `json:"created_at"`
}
