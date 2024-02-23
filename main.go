package main

import (
	"chatapp/pkg/routes"
	"github.com/gin-gonic/gin"
	"os"
	"chatapp/pkg/utils"
)

func main() {
    utils.LoadEnv()
	r := gin.Default()
	// get routers
	routes.Routers(r)
	// serve default port 8080
	port := os.Getenv("PORT")
	err := r.Run("0.0.0.0:" + port)
	if err != nil {
		panic(err)
	}
}
