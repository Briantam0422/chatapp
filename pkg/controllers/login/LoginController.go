package login

import (
	"chatapp/pkg/models"
	"chatapp/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var u models.UserRequest
	// get request data
	err := c.ShouldBind(&u)
	if err != nil {
		utils.ErrorRespond(c, err)
		return
	}

	// check has username
	hasUser := models.HasUser(u.Username)
	if !hasUser {
		utils.ErrorRespondWithMessage(c, "User not found")
		return
	}
	// check password
	_, user := models.FindUserByUsername(u.Username)
	authorized := user.CheckPasswordHash(u.Password)
	if !authorized {
		utils.ErrorRespondWithMessage(c, "Invalid Username and Password")
		return
	}
	// create jwt token
	expirationTime := time.Now().Add(60 * time.Minute)
	tokenString, err := user.GenerateToken(expirationTime)
	if err != nil {
		utils.ErrorRespond(c, err)
		return
	}

	// set browser cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Path: "/",
	})

	// return json respond
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"id":       u.Id,
		"username": u.Username,
		"token":    tokenString,
	})
}

func IsAuth(c *gin.Context) {
	isAuth := c.MustGet("isAuth").(bool)
	if !isAuth {
		utils.UnAuthorized(c, "Please Login First")
		return
	}
	// get cookie token
	token, err := c.Cookie("token")
	fmt.Println(token)
	if err != nil {
		log.Println(err)
		c.Set("isAuth", false)
		return
	}
	r, u := models.FindUserByToken(token)
	fmt.Println(r.Row())
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"id":       u.Id,
		"username": u.Username,
	})
}
