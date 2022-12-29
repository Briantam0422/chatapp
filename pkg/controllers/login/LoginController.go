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
	utils.CheckError(err)
	c.JSON(http.StatusOK, gin.H{
		"status":   "posted",
		"username": u.Username,
		"password": u.Password,
	})
}
