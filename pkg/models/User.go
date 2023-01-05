package models

import (
	"chatapp/pkg/utils"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
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

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func HasUser(username string) bool {
	result, _ := FindUserByUsername(username)
	fmt.Println("result: ", result.RowsAffected)
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

func FindUserByUsername(username string) (*gorm.DB, User) {
	u := User{
		Username: username,
	}
	db := utils.ConnectDB()
	defer utils.CloseDB(db)
	fmt.Println(u.Username)
	result := db.Find(&u, "username = ?", u.Username)
	return result, u
}

func (u User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u User) GenerateToken(expirationTime time.Time) (string, error) {
	// expiry time
	claims := &Claims{
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	utils.LoadEnv()
	jwtKey := os.Getenv("JWT_KEY")
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	// save token to db
	db := utils.ConnectDB()
	defer utils.CloseDB(db)

	db.Model(&u).Where("id = ?", u.Id).Update("token", tokenString)

	return tokenString, nil
}

func CreateUser(u UserRequest) error {
	db := utils.ConnectDB()
	defer utils.CloseDB(db)
	hasUser := HasUser(u.Username)
	if hasUser {
		return errors.New("user has already existed")
	}
	hash, err := utils.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	newUser := User{
		Username: u.Username,
		Password: u.Password,
	}
	createdUser := db.Create(&newUser)
	err = createdUser.Error
	if err != nil {
		return err
	}
	return nil
}
