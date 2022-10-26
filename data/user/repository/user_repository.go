package repository

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/VooDooStack/FitStackAPI/domain/dto"
	domain "github.com/VooDooStack/FitStackAPI/domain/user"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepository struct {
	Database         gorm.DB
	FriendRepository domain.FriendshipRepository
}

func NewUserRepository(db gorm.DB, fr domain.FriendshipRepository) domain.UserRepository {
	return &userRepository{db, fr}
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

func (u *userRepository) CheckUniqueFields(user *dto.UserSignUp) error {
	// var userCheck *domain.User
	// sqlStatement := fmt.Sprintf(`
	// SELECT * FROM users WHERE email='%s'
	// OR (display_name='%s')
	// OR (phone_number='%s')
	// `, user.Email, user.DisplayName, user.PhoneNumber)

	// err := u.Database.QueryRow(context.Background(), sqlStatement).Scan(&userCheck)
	// if err == nil {
	// 	logrus.Error(err)
	// 	return fmt.Errorf("user already exists")
	// }

	return nil
}

func (u *userRepository) UpdateUserAvatar(uuid string, fileURL string) error {
	// sqlStatement := `
	// UPDATE user_profiles SET avatar = $1
	// WHERE user_profiles.id = $2;
	// `
	// _, err := u.Database.Exec(context.Background(), sqlStatement, fileURL, uuid)
	// if err != nil {
	// 	logrus.Error("error querying database error: %v", err)
	// 	return fmt.Errorf("error updating user record: %v", err)
	// }

	return nil
}

func (u *userRepository) GetUserProfile(uuid string) (*domain.UserProfile, error) {
	profile := domain.UserProfile{}
	// sqlStatement := `
	// SELECT * FROM user_profiles
	// WHERE id= $1;
	// `
	// row, err := u.Database.Query(context.Background(), sqlStatement, uuid)
	// if err != nil {
	// 	logrus.Error("error querying database error: %v", err)
	// 	return nil, err
	// }

	// err = pgxscan.ScanOne(&profile, row)
	// if err != nil {
	// 	logrus.Error("error scanning row error: %v", err)
	// 	return nil, err
	// }

	// friends, err := u.FriendRepository.GetFriends(uuid)
	// if err != nil {
	// 	logrus.Error("error getting friends: %v", err)
	// 	return nil, err
	// }

	// profile.Friends = friends

	return &profile, nil
}
