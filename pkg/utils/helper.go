package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorRespond(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "failed",
		"message": "Server Error",
		"err":     err.Error(),
	})
}

func Load(path string) {
	err := godotenv.Load("../main/.env")
	CheckError(err)
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Hash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 10)
	return string(bytes), err
}
