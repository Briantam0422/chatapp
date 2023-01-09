package chat

import (
	"fmt"
)

type Rooms struct {
	rooms      map[int32]*Room
	register   chan *Room
	unregister chan *Room
}

func Initialize() *Rooms {
	rooms := &Rooms{
		rooms:      map[int32]*Room{},
		register:   make(chan *Room),
		unregister: make(chan *Room),
	}
	go rooms.run()
	return rooms
}

func (rs *Rooms) run() {
	for {
		select {
		case room := <-rs.register:
			fmt.Println("room registered... room id -", room.id)
			rs.rooms[room.id] = room
			fmt.Println("rooms", len(rs.rooms))
		case room := <-rs.unregister:
			if _, ok := rs.rooms[room.id]; ok {
				delete(rs.rooms, room.id)
				for client := range room.clients {
					close(client.send)
					fmt.Println("clients unregistered", len(room.clients))
				}
				fmt.Println("room unregistered", len(rs.rooms))
			}
		}
	}
}
