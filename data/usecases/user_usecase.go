package usecases

import (
	"fmt"

	"github.com/VooDooStack/FitStackAPI/domain"
	"gorm.io/gorm"
)

func GetUserByEmail(db *gorm.DB, email string) (domain.User, error) {
	var user domain.User

	userCheck := db.Where(domain.User{Email: email}).Find(&user)
	if userCheck.Error != nil {
		fmt.Print(userCheck.Error)
		return domain.User{DisplayName: "Null User"}, userCheck.Error
	}

	return user, nil
}

func GetUserByUuid(db *gorm.DB, uuid string) (domain.User, error) {
	var user domain.User

	userCheck := db.Where(domain.User{Uuid: uuid}).Find(&user)
	if userCheck.Error != nil {
		fmt.Print(userCheck.Error)
		return domain.User{DisplayName: "Null User"}, userCheck.Error
	}

	return user, nil
}

func GetUserByDisplayName(db *gorm.DB, displayName string) (domain.User, error) {
	var user domain.User

	userCheck := db.Where(domain.User{DisplayName: displayName}).Find(&user)
	if userCheck.Error != nil {
		fmt.Print(userCheck.Error)
		return domain.User{DisplayName: "Null User"}, userCheck.Error
	}

	return user, nil
}

func CreateUser(db *gorm.DB, user domain.User) (domain.User, error) {
	var newUser domain.User

	if result := db.Create(&user); result.Error != nil {
		fmt.Print(result.Error)
		return domain.User{DisplayName: "Null User"}, result.Error
	}

	return newUser, nil
}

func UpdateUser(db *gorm.DB, user domain.User) (domain.User, error) {
	//TODO: Implement
	var updatedUser domain.User

	if result := db.Save(&user).Where(domain.User{DisplayName: user.DisplayName}); result.Error != nil {
		fmt.Print(result.Error)
		return domain.User{DisplayName: "Null User"}, result.Error
	}

	return updatedUser, nil
}

func DeleteUserByEmail(db *gorm.DB, email string) (domain.User, error) {
	//TODO: Implement
	var user domain.User

	if result := db.Where(domain.User{Email: email}).Delete(&user); result.Error != nil {
		fmt.Print(result.Error)
		return domain.User{DisplayName: "Null User"}, result.Error
	}

	return user, nil
}

func DeleteUserByUuid(db *gorm.DB, uuid string) (domain.User, error) {
	//TODO: Implement
	var user domain.User

	if result := db.Where(domain.User{Uuid: uuid}).Delete(&user); result.Error != nil {
		fmt.Print(result.Error)
		return domain.User{DisplayName: "Null User"}, result.Error
	}

	return user, nil
}
