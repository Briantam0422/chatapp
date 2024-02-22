package middlewares

import (
	"chatapp/pkg/models"
// 	"chatapp/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get cookie token
		token, err := c.Cookie("token")
		log.Println(token)
		fmt.Println(token)
		if err != nil {
			log.Println(err)
			c.Set("isAuth", false)
			return
		}

		//get env jwtKey
// 		utils.LoadEnv()
		jwtKey := os.Getenv("JWT_KEY")

		// initialize a new instance of claims
		claims := &models.Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		if err != nil {
			c.Set("isAuth", false)
			log.Println(err)
			return
		}

		// check token validation
		if !tkn.Valid {
			c.Set("isAuth", false)
			log.Println("Token is invalid")
			return
		}

		c.Set("isAuth", true)
	}
}
