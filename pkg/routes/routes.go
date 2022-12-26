package routes

import (
	"chatapp/pkg/controllers/login"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	r.POST("/", login.Login)
}
