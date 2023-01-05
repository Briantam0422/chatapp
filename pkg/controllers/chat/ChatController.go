package chat

import (
	"chatapp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func Initial(c *gin.Context) {
	isAuth := c.MustGet("isAuth").(bool)
	if !isAuth {
		utils.UnAuthorized(c, "Please Login First")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func Connect(c *gin.Context) {
	u := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, _ := u.Upgrade(c.Writer, c.Request, nil)

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		serverMessage := "Server Message: " + string(p)
		if err := conn.WriteMessage(messageType, []byte(serverMessage)); err != nil {
			log.Println(err)
			return
		}
	}
}
