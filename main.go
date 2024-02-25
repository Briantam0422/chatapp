package main

import (
	"chatapp/pkg/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	//utils.LoadEnv()
	r := gin.Default()
	// get routers
	routes.Routers(r)
	port := os.Getenv("PORT")
	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
