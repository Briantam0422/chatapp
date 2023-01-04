package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token := c.Query("token")

		c.Set("isAuth", true)
	}
}
