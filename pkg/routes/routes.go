package routes

import (
	"chatapp/pkg/controllers/chat"
	"chatapp/pkg/controllers/login"
	"chatapp/pkg/controllers/registration"
	"chatapp/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	r.POST("/login", login.Login)
	r.POST("/register", registration.Registration)
	r.Use(middlewares.AuthRequired())
	r.Group("chat")
	{
		r.GET("/initial", chat.Initial)
		r.GET("/connect", chat.Connect)
	}
}
