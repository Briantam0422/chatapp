package login

import (
	"chatapp/pkg/models"
	"chatapp/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var u models.UserRequest
	// get request data
	err := c.ShouldBind(&u)
	utils.ErrorRespond(c, err)

	// check has username
	hasUser := models.HasUser(u.Username)
	if !hasUser {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "User not found",
		})
		return
	}
	// check password
	_, user := models.FindUserByUsername(u.Username)
	authorized := user.CheckPasswordHash(u.Password)
	if !authorized {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Invalid Username and Password",
		})
		return
	}
	// create jwt token
	tokenString, err := user.GenerateToken()
	utils.ErrorRespond(c, err)
	c.SetCookie("token", tokenString, 9999, "/", "localhost", true, false)

	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"username": u.Username,
		"token":    tokenString,
	})
}
