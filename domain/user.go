package domain

type User struct {
	UserId        string       `gorm:"primaryKey" json:"user_id" binding:"required"`
	Email         string       `gorm:"unique;not null" json:"email" binding:"required,email"`
	DisplayName   string       `gorm:"unique;not null" json:"display_name" binding:"required"`
	FirstName     string       `json:"first_name" binding:"required"`
	LastName      string       `json:"last_name" binding:"required"`
	PhoneNumber   string       `gorm:"unique;not null" json:"phone_number" binding:"required"`
	DateOfBirth   string       `json:"date_of_birth" binding:"required"`
	EmailVerified bool         `json:"email_verified"`
	PhotoURL      string       `json:"photo_url" binding:"omitempty,url"`
	Friendship    []Friendship `gorm:"many2many:user_friendships" json:"friendship"`
	UpdatedAt     string       `json:"updated_at"`
	CreatedAt     int64        `json:"created_at"`
	RefreshToken  string       `json:"refresh_token"`
}

type UserUsecase interface {
	SignUp(user User) (User, error)
	GetByUuid(uuid string) (User, error)
	Update(uuid string) error
	GetByEmail(email string) (User, error)
	Store(user User) error
	Delete(uuid string) error
}

type UserRepository interface {
	SignUp(user User) (User, error)
	GetByUuid(uuid string) (User, error)
	Update(uuid string) error
	GetByEmail(email string) (User, error)
	Store(user User) error
	Delete(uuid string) error
}
