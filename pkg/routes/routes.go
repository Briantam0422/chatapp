package routes

import (
	"chatapp/pkg/controllers/chat"
	"chatapp/pkg/controllers/login"
	"chatapp/pkg/controllers/registration"
	"chatapp/pkg/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Routers(r *gin.Engine) {
	// cross origins
	//r.Use(cors.Default())
	// list font end localhost port 5173
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	// API
	r.POST("/login", login.Login)
	r.POST("/register", registration.Registration)
	// Auth Middlewares
	r.Use(middlewares.AuthRequired())
	r.GET("/isAuth", login.IsAuth)
	// Chat API
	r.Group("chat")
	{
		rooms := chat.Initialize()
		r.GET("/start", chat.Chat(rooms))
		r.GET("/initial", chat.Initial)
		r.GET("/new", chat.New(rooms))
		r.GET("/close", chat.Close(rooms))
	}
}
