package models

import (
	"chatapp/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        int            `json:"id,identity"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HasUser(username string) bool {
	result := FindUserByUsername(username)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

func FindUserById(id int) *gorm.DB {
	db := utils.ConnectDB()
	defer utils.CloseDB(db)
	result := db.First(&User{Id: id})
	return result
}

func FindUserByUsername(username string) *gorm.DB {
	db := utils.ConnectDB()
	defer utils.CloseDB(db)
	result := db.First(&User{Username: username})
	return result
}

func (u *User) CheckPasswordHash(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.Password))
	return err == nil
}

func CreateUser(u UserRequest) (string, error) {
	db := utils.ConnectDB()
	defer utils.CloseDB(db)
	hasUser := HasUser(u.Username)
	if hasUser {
		return "User has already existed", nil
	}
	hash, err := utils.Hash(u.Password)
	if err != nil {
		return "Server Error", err
	}
	u.Password = hash
	newUser := User{
		Username: u.Username,
		Password: u.Password,
	}
	createdUser := db.Create(&newUser)
	err = createdUser.Error
	if err != nil {
		return "Server Error", err
	}
	return "", nil
}
