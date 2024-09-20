package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	println("hello world")

	app := fiber.New()
	log.Fatal(app.Listen(":8080"))
}
