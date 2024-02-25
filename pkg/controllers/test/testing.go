package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestingAPI(c *gin.Context) {
	// return json respond
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
