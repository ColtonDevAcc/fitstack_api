package domain

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// UserUsecase represent the users's use cases
type UserUseCase interface {
	Fetch(ctx gin.Context, cursor string, num int64) ([]User, string, error)
	GetByID(ctx gin.Context, id int64) (User, error)
	Update(ctx gin.Context, ar *User) error
	GetByTitle(ctx gin.Context, title string) (User, error)
	Store(gin.Context, *User) error
	Delete(ctx gin.Context, id int64) error
}

// UserRepository represent the users's repository contract
type UserRepository interface {
	Fetch(ctx gin.Context, cursor string, num int64) (res []User, nextCursor string, err error)
	GetByID(ctx gin.Context, id int64) (User, error)
	GetByTitle(ctx gin.Context, title string) (User, error)
	Update(ctx gin.Context, ar *User) error
	Store(ctx gin.Context, a *User) error
	Delete(ctx gin.Context, id int64) error
}
