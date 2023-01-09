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
			fmt.Println("room id covert err")
			return
		}
		r := rooms.rooms[int32(rId)]
		if r == nil {
			fmt.Println("room doesn't exist")
			return
		}
		fmt.Println("client id : ", id)
		fmt.Println("enter room : ", roomId)
		client := newClient(id, r, c.Writer, c.Request)
		client.room.register <- client
		fmt.Println(client.room)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
	return gin.HandlerFunc(fn)
}

func Close(rooms *Rooms) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		roomId := c.Query("room_id")
		rId, err := strconv.ParseInt(roomId, 10, 64)
		if err != nil {
			fmt.Println("room id covert err")
			return
		}
		room := rooms.rooms[int32(rId)]
		if room == nil {
			fmt.Println("room doesn't exist")
			return
		}
		rooms.unregister <- room
		fmt.Println("room has been closed")
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
	return gin.HandlerFunc(fn)
}
