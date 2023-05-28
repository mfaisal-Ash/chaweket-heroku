package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/rofinafiin/ws-repo/module"
	"github.com/rofinafiin/ws-repo/typestruct"
	"log"
)

func Websocket(c *websocket.Conn) {
	username := c.Query("username")
	client := &typestruct.Client{
		Username: username,
		Conn:     c,
	}
	module.NewChatRoom().Register <- client

	defer func() {
		module.NewChatRoom().Unregister <- client
		c.Close()
	}()

	for {
		var message typestruct.Message
		err := c.ReadJSON(&message)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		module.BroadcastMessage(message)
	}
}

func GetIP(c *fiber.Ctx) error {
	getip := musik.GetIPaddress()
	return c.JSON(getip)
}
