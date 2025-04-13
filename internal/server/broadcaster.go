package server

import (
	"log"

	"github.com/gorilla/websocket"
)

var Clients = make(map[*websocket.Conn]bool)
var Broadcast = make(chan string)

// StartBroadcastLoop starts the message broadcaster
func StartBroadcastLoop() {
	go func() {
		for {
			msg := <-Broadcast
			for client := range Clients {
				err := client.WriteMessage(websocket.TextMessage, []byte(msg))
				if err != nil {
					log.Println("Write error:", err)
					client.Close()
					delete(Clients, client)
				}
			}
		}
	}()
}
