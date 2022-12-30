package login

import (
	"chatapp/pkg/models"
	"chatapp/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var u models.UserRequest
	err := c.ShouldBind(&u)
	utils.ErrorRespond(c, err)
	hasUser := models.HasUser(u.Username)
	if !hasUser {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "User not found",
		})
		return
	}
	_, user := models.FindUserByUsername(u.Username)
	authorized := user.CheckPasswordHash(u.Password)
	if !authorized {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Invalid Username and Password",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"username": u.Username,
	})
}
