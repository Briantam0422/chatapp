package login

import (
	"chatapp/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var u user
	err := c.ShouldBind(&u)
	utils.CheckError(err)
	c.JSON(http.StatusOK, gin.H{
		"status":   "posted",
		"username": u.Username,
		"password": u.Password,
	})
}
