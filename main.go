package main

import (
	"log"

	"github.com/bike-sharing-app/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("api/bikes", handlers.GetAllBikes)

	log.Fatal(app.Listen(":3000"))
}
