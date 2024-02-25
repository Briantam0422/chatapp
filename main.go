package main

import (
	"chatapp/pkg/routes"
	"chatapp/pkg/utils"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	utils.LoadEnv()
	r := gin.Default()
	// get routers
	routes.Routers(r)
	port := os.Getenv("PORT")
	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
