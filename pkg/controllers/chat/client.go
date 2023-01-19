package chat

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	id   string
	name string
	room *Room
	conn *websocket.Conn
	send chan *Message
}

func (c *Client) receiveMessages() {
	defer func() {
		c.room.unregister <- c
		err := c.conn.Close()
		if err != nil {
			fmt.Println("Close connection failed")
		}
	}()
	c.conn.SetReadLimit(maxMessageSize)
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		fmt.Println("read dead line timeout")
	}
	c.conn.SetPongHandler(func(gg string) error {
		fmt.Println("pong hit", gg)
		err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			fmt.Println("read dead line timeout")
		}
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		t := time.Now()
		c.room.broadcast <- &Message{
			Message:  string(message),
			ClientID: c.id,
			Type:     "text",
			Username: c.name,
			Time:     strconv.Itoa(t.Year()) + "-" + strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(int(t.Day())) + " " + strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute()),
		}
	}
}

func (c *Client) sendMessages() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		err := c.conn.Close()
		if err != nil {
			fmt.Println("connection close error")
		}
	}()

	for {
		select {
		case message, ok := <-c.send:
			// sets write timeout of 10 seconds
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.conn.WriteJSON(message)
			if err != nil {
				return
			}

		// this sends a ping to the connection very 54 seconds
		case <-ticker.C:
			fmt.Println("ticker hit")
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func newClient(id string, name string, room *Room, w http.ResponseWriter, r *http.Request) *Client {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	client := &Client{id: id, name: name, room: room, conn: conn, send: make(chan *Message, 256)}

	go client.sendMessages()
	go client.receiveMessages()

	return client
}
