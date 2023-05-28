package struct

import "github.com/gofiber/websocket/v2"

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type Client struct {
	Username string
	Conn     *websocket.Conn
}

type ChatRoom struct {
	Clients    []*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
}

type NewMessage struct {
	Id      string
	Message string
}
