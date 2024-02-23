package routes

import (
	"chatapp/pkg/controllers/chat"
	"chatapp/pkg/controllers/login"
	"chatapp/pkg/controllers/registration"
	"chatapp/pkg/middlewares"
// 	"chatapp/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func Routers(r *gin.Engine) {
	// cross origins
	// listen font end localhost port 5173
// 	utils.LoadEnv()
	productionOrigin := os.Getenv("APP_URL")
	devOrigin := os.Getenv("APP_URL") + ":" + os.Getenv("DEV_FRONTEND_PORT")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{productionOrigin, devOrigin, "http://0.0.0.0", "https://0.0.0.0", "https://chatapp.api.briantambusiness.com"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
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
