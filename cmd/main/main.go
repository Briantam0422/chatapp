package main

import (
	"chatapp/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// get routers
	routes.Routers(r)
	// serve default port 8080
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
