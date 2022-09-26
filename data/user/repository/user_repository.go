package repository

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepository struct {
	Database gorm.DB
}

func NewUserRepository(db gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (u *userRepository) Delete(uuid string) error {
	tx := u.Database.Where(domain.User{UserId: uuid}).Delete(&domain.User{})
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return tx.Error
	}

	return nil
}

func (u *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	tx := u.Database.Where(domain.User{Email: email}).Find(&user)
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return domain.User{DisplayName: "Null User"}, tx.Error
	}

	return user, nil
}

func (u *userRepository) GetByUuid(uuid string) (domain.User, error) {
	var user domain.User
	tx := u.Database.Where(domain.User{UserId: uuid}).Find(&user)
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return domain.User{DisplayName: "Null User"}, tx.Error
	}

	return user, nil
}

func (u *userRepository) Store(user domain.User) error {
	tx := u.Database.Create(&user)
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return tx.Error
	}

	return nil
}

func (u *userRepository) Update(uuid string) error {
	tx := u.Database.Where(domain.User{UserId: uuid}).Save(&domain.User{})
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return tx.Error
	}

	return nil
}

func (u *userRepository) SignUp(user domain.User) (domain.User, error) {
	tx := u.Database.Create(&user)
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return domain.User{DisplayName: "Null User"}, tx.Error
	}

	return user, nil
}

func (u *userRepository) SignInWithToken(uuid string) (domain.User, error) {
	user, err := u.GetByUuid(uuid)
	if err != nil {
		logrus.Error(err)

		return domain.User{}, err
	}
	// u.Database.Where(user).Save()
	return user, nil
}

func (u *userRepository) RefreshToken(refresh_token string) (string, error) {
	fmt.Println("exchanging token " + refresh_token + "/n")

	jsonBody := []byte(fmt.Sprintf(refresh_token, "grant_type=refresh_token&refresh_token=%d"))
	bodyReader := bytes.NewReader(jsonBody)

	requestURL := "https://securetoken.googleapis.com/v1/token?key=AIzaSyD00CGxTDoVgQEKBKLh2xIfudyaIHifS5Y"
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb := strings.Split(string(body), "%")[0]

	fmt.Println("new token response " + sb)

	return sb, nil
}
