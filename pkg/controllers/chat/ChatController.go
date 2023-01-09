package chat

import (
	"chatapp/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func New(rooms *Rooms) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		roomName := c.Query("room_name")
		room := NewRoom(rooms)
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"room_name": roomName,
			"room_id":   room.id,
		})
	}
	return gin.HandlerFunc(fn)
}

func Chat(rooms *Rooms) gin.HandlerFunc {
	//room := newRoom()
	fn := func(c *gin.Context) {
		id := c.Query("id")
		roomId := c.Query("room_id")
		rId, err := strconv.ParseInt(roomId, 10, 64)
		if err != nil {
			panic(err)
		}
		r := rooms.rooms[int32(rId)]
		fmt.Println("client id : ", id)
		fmt.Println("enter room : ", roomId)
		client := newClient(id, r, c.Writer, c.Request)
		client.room.register <- client
		fmt.Println(client.room)
	}
	return gin.HandlerFunc(fn)

}

//func Connect(c *gin.Context) {
//	u := websocket.Upgrader{
//		ReadBufferSize:  1024,
//		WriteBufferSize: 1024,
//	}
//	conn, _ := u.Upgrade(c.Writer, c.Request, nil)
//	// receive user info
//	messageType, p, err := conn.ReadMessage()
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	// return user info
//	serverMessage := "Server Message: " + string(p)
//	if err := conn.WriteMessage(messageType, []byte(serverMessage)); err != nil {
//		log.Println(err)
//		return
//	}
//
//}
