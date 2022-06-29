package models

type User struct {
	ID        int    `json:"id"`
	Uuid      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	Token     string `json:"token"`
}
