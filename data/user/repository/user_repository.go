package repository

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/VooDooStack/FitStackAPI/domain/dto"
	"github.com/goccy/go-json"
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
	tx := u.Database.Where(domain.User{UUID: uuid}).Delete(&domain.User{})
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
	tx := u.Database.Where(domain.User{UUID: uuid}).Find(&user)
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
	tx := u.Database.Where(domain.User{UUID: uuid}).Save(&domain.User{})
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
	jsonBody := []byte(fmt.Sprintf(refresh_token, "grant_type=refresh_token&refresh_token=%d"))
	bodyReader := bytes.NewReader(jsonBody)

	key := os.Getenv("API_KEY")
	requestURL := fmt.Sprintf("https://securetoken.googleapis.com/v1/token?key=%q", key)
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		logrus.Error(err)
	}

	//We Read the response body on the line below.
	body, err := io.ReadAll(req.Body)
	if err != nil {
		logrus.Error(err)
	}

	//Convert the body to type string
	sb := strings.Split(string(body), "%")[0]

	return sb, nil
}

func (u *userRepository) SignInWithEmailAndPassword(login *dto.LoginInEmailAndPassword) (string, error) {
	data, err := json.Marshal(&login)
	if err != nil {
		logrus.Error(err)
	}
	bodyReader := bytes.NewReader(data)

	key := os.Getenv("API_KEY")
	requestURL := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%q", key)
	res, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		logrus.Error(err)
	}

	//We Read the response body on the line below.
	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Error(err)
	}

	//Convert the body to type string
	return string(body), nil
}

func (u *userRepository) CheckUniqueFields(user domain.User) error {
	var userEmailCheck domain.User
	tx := u.Database.Where(domain.User{Email: user.Email}).Take(&userEmailCheck)
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return tx.Error
	}

	var userDisplayNameCheck domain.User
	tx = u.Database.Where(domain.User{DisplayName: user.DisplayName}).Take(&userDisplayNameCheck)
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return tx.Error
	}

	var userPhoneNumberCheck domain.User
	tx = u.Database.Where(domain.User{PhoneNumber: user.PhoneNumber}).Take(&userPhoneNumberCheck)
	if tx.Error != nil {
		logrus.Error(tx.Error)

		return tx.Error
	}
	emptyUser := domain.User{}
	if userDisplayNameCheck.DisplayName != emptyUser.DisplayName && userPhoneNumberCheck.PhoneNumber != emptyUser.PhoneNumber && userEmailCheck.Email != emptyUser.Email {
		return fmt.Errorf("user already exits")
	}

	return nil
}
