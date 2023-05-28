package main

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/sidiq200/chaweket-heroku/module"
	"github.com/sidiq200/chaweket-heroku/url"
	"log"
)

func main() {
	go module.Run()

	app := fiber.New()
	url.Web(app)
	log.Fatal(app.Listen(musik.Dangdut()))
}
