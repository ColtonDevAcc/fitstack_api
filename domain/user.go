package domain

type User struct {
	ID            int    `json:"id"`
	Uuid          string `json:"uuid"`
	DisplayName   string `json:"display_name"`
	PhoneNumber   string `json:"phone_number"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	PhotoURL      string `json:"photo_url"`
	UpdatedAt     string `json:"updated_at"`
	CreatedAt     int64  `json:"created_at"`
	RefreshToken  string `json:"refresh_token"`
}

type UserUsecase interface {
	GetByUuid(uuid string) (User, error)
	Update(uuid string) error
	GetByEmail(email string) (User, error)
	Store(user User) error
	Delete(uuid string) error
}

type UserRepository interface {
	GetByUuid(uuid string) (User, error)
	Update(uuid string) error
	GetByEmail(email string) (User, error)
	Store(user User) error
	Delete(uuid string) error
}
