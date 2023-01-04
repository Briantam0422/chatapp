package registration

import (
	"chatapp/pkg/models"
	"chatapp/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Registration(c *gin.Context) {
	var u models.UserRequest
	err := c.ShouldBind(&u)
	if err != nil {
		utils.ErrorRespond(c, err)
		return
	}
	err = models.CreateUser(u)
	if err != nil {
		utils.ErrorRespond(c, err)
		return
	}
	// create jwt token
	result, user := models.FindUserByUsername(u.Username)
	err = result.Error
	if err != nil {
		utils.ErrorRespond(c, err)
		return
	}
	tokenString, err := user.GenerateToken()
	if err != nil {
		utils.ErrorRespond(c, err)
		return
	}
	c.SetCookie("token", tokenString, 9999, "/", "localhost", true, false)

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"username": u.Username,
		"token":    tokenString,
	})

}
