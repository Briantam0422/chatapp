package main

import (
	"chatapp/pkg/routes"
	"github.com/gin-gonic/gin"
	"chatapp/pkg/utils"
)

func main() {
    utils.LoadEnv()
	r := gin.Default()
	// get routers
	routes.Routers(r)
	// serve default port 8080
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
