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
	utils.ErrorRespond(c, err)
	res, err := models.CreateUser(u)
	if res != "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": res,
			"err":     err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"username": u.Username,
	})
}
