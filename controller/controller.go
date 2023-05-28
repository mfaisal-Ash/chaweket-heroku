package controller

import (
	"fmt"
    "github.com/gorilla/websocket"
    "github.com/novalagung/gubrak/v2"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
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
