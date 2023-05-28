package module

import (
	"fmt"
	"github.com/rofinafiin/ws-repo/typestruct"
	"log"
)

func NewChatRoom() *typestruct.ChatRoom {
	return &typestruct.ChatRoom{
		Clients:    make([]*typestruct.Client, 0),
		Register:   make(chan *typestruct.Client),
		Unregister: make(chan *typestruct.Client),
		Broadcast:  make(chan typestruct.Message),
	}
}
func BroadcastMessage(message typestruct.Message) {
	NewChatRoom().Broadcast <- message
}

func Run() {
	for {
		select {
		case client := <-NewChatRoom().Register:
			NewChatRoom().Clients = append(NewChatRoom().Clients, client)
			go BroadcastMessage(typestruct.Message{
				Username: "Server",
				Content:  fmt.Sprintf("User %s joined the chat", client.Username),
			})
		case client := <-NewChatRoom().Unregister:
			for i, c := range NewChatRoom().Clients {
				if c == client {
					NewChatRoom().Clients = append(NewChatRoom().Clients[:i], NewChatRoom().Clients[i+1:]...)
					go BroadcastMessage(typestruct.Message{
						Username: "Server",
						Content:  fmt.Sprintf("User %s left the chat", client.Username),
					})
					break
				}
			}
		case message := <-NewChatRoom().Broadcast:
			for _, client := range NewChatRoom().Clients {
				go func(c *typestruct.Client) {
					if err := c.Conn.WriteJSON(message); err != nil {
						log.Println("Error broadcasting message:", err)
					}
				}(client)
			}
		}
	}
}
