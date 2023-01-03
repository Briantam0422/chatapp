package routes

import (
	"chatapp/pkg/controllers/chat"
	"chatapp/pkg/controllers/login"
	"chatapp/pkg/controllers/registration"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	r.GET("/", login.Login)
	r.POST("/login", login.Login)
	r.POST("/register", registration.Registration)
	r.GET("/chat/connect", chat.Connect)
}
