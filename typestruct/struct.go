package typestruct

import "github.com/gofiber/websocket/v2"

type WebSocketConnection struct {
	*websocket.Conn
	Username string
}

type SocketPayload struct {
	Message string
}

type SocketResponse struct {
	From    string
	Type    string
	Message string
}

type Client struct {
	Username string
	Conn     *websocket.Conn
}

type NewMessage struct {
	Id      string
	Message string
}
