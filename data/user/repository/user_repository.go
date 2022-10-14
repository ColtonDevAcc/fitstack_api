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
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	Database pgxpool.Pool
}

func NewUserRepository(db pgxpool.Pool) domain.UserRepository {
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
	user := domain.User{}
	sqlStatement := `
	SELECT * FROM users
	WHERE id=$1;
	`
	row, err := u.Database.Query(context.Background(), sqlStatement, uuid)
	if err != nil {
		logrus.Error("error querying database error: %v", err)
		return nil, err
	}

	err = pgxscan.ScanOne(&user, row)
	if err != nil {
		logrus.Error("error scanning row error: %v", err)
		return nil, err
	}

	return &user, nil
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

func (u *userRepository) SignUp(user *dto.UserSignUp) (*domain.User, error) {
	newUser := domain.User{}
	sqlStatement := `
	INSERT INTO users (id, display_name, first_name, last_name, phone_number, phone_verified, date_of_birth, email, email_verified, photo_url)
	VALUES ($1, $2, $3, $4, $5, $6, $7 ,$8, $9, $10)
	RETURNING *
	`

	rows, err := u.Database.Query(context.Background(), sqlStatement, &user.Id, &user.DisplayName, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.PhoneVerified, &user.DateOfBirth, &user.Email, &user.EmailVerified, &user.PhotoURL)
	if err != nil {
		logrus.Error(fmt.Printf("error querying row err: %v", err))
		return nil, err
	}

	defer rows.Close()
	err = rows.Scan(&newUser)
	if err != nil {
		logrus.Error(fmt.Printf("error scanning row err: %v", err))
		return nil, err
	}

	return &newUser, nil
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

func (u *userRepository) CheckUniqueFields(user *dto.UserSignUp) error {
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

func (u *userRepository) UpdateUserAvatar(uuid string, fileURL string) error {
	sqlStatement := `
	UPDATE users SET photo_url = $1
	WHERE users.id = $2;
	`
	_, err := u.Database.Exec(context.Background(), sqlStatement, fileURL, uuid)
	if err != nil {
		logrus.Error("error querying database error: %v", err)
		return fmt.Errorf("error updating user record: %v", err)
	}

	return nil
}
