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

func UnAuthorized(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "unauthorized",
		"message": message,
	})
}

func ErrorRespond(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "error",
		"message": err.Error(),
		"err":     "Server Error",
	})
}

func ErrorRespondWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "error",
		"message": message,
	})
}

func Load(path string) {
	err := godotenv.Load(".env")
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
