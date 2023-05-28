package main

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/rofinafiin/ws-repo/module"
	"github.com/rofinafiin/ws-repo/url"
	"log"
)

func main() {
	go module.Run()

	app := fiber.New()
	url.Web(app)
	log.Fatal(app.Listen(musik.Dangdut()))
}
