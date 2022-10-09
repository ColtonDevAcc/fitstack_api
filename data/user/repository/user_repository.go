package repository

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/VooDooStack/FitStackAPI/domain"
	"github.com/VooDooStack/FitStackAPI/domain/dto"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/goccy/go-json"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	Database pgx.Conn
}

func NewUserRepository(db pgx.Conn) domain.UserRepository {
	return &userRepository{db}
}

func (u *userRepository) Delete(uuid string) error {
	//TODO:
	// err := u.Database.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	// if err != nil {
	// 	logrus.Error(err)

	// 	return err
	// }

	return nil
}

func (u *userRepository) GetByEmail(email string) (*domain.User, error) {
	//TODO:
	// var user domain.User
	// tx := u.Database.Where(domain.User{Email: email}).Find(&user)
	// if tx.Error != nil {
	// 	logrus.Error(tx.Error)

	// 	return domain.User{DisplayName: "Null User"}, tx.Error
	// }

	return nil, nil
}

func (u *userRepository) GetByUuid(uuid string) (*domain.User, error) {
	// var user domain.User
	// tx := u.Database.Where(domain.User{UUID: uuid}).Find(&user)
	// if tx.Error != nil {
	// 	logrus.Error(tx.Error)

	// 	return domain.User{DisplayName: "Null User"}, tx.Error
	// }

	return nil, nil
}

func (u *userRepository) Store(user *domain.User) error {
	// tx := u.Database.Create(&user)
	// if tx.Error != nil {
	// 	logrus.Error(tx.Error)

	// 	return tx.Error
	// }

	return nil
}

func (u *userRepository) Update(uuid string) error {
	// tx := u.Database.Where(domain.User{UUID: uuid}).Save(&domain.User{})
	// if tx.Error != nil {
	// 	logrus.Error(tx.Error)

	// 	return tx.Error
	// }

	return nil
}

func (u *userRepository) SignUp(user *domain.User) (*domain.User, error) {
	sqlStatement := `
	INSERT INTO users (id, display_name, first_name, last_name, phone_number, phone_verified, date_of_birth, email, email_verified)
	VALUES ($1, $2, $3, $4, $5, $6, $7 ,$8, $9)
	RETURNING *
	`

	rows, _ := u.Database.Query(context.Background(), sqlStatement, &user.Id, &user.DisplayName, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.PhoneVerified, &user.DateOfBirth, &user.Email, &user.EmailVerified)
	defer rows.Close()

	pgxscan.ScanRow(&user, rows)

	return user, nil
}

func (u *userRepository) SignInWithToken(uuid string) (*domain.User, error) {
	// user, err := u.GetByUuid(uuid)
	// if err != nil {
	// 	logrus.Error(err)

	// 	return domain.User{}, err
	// }

	// u.Database.Where(user).Save()
	return nil, nil
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

func (u *userRepository) CheckUniqueFields(user *domain.User) error {
	var userCheck *domain.User
	sqlStatement := fmt.Sprintf(`
	SELECT * FROM users WHERE email='%s'
	OR (display_name='%s')
	OR (phone_number='%s')
	`, user.Email, user.DisplayName, user.PhoneNumber)

	err := u.Database.QueryRow(context.Background(), sqlStatement).Scan(&userCheck)
	if err == nil {
		logrus.Error(err)
		return fmt.Errorf("user already exists")
	}

	return nil
}
