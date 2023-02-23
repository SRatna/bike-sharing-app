package main

import (
	"log"

	"github.com/bike-sharing-app/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./dist")

	app.Get("api/bikes", handlers.GetAllBikes)
	app.Patch("api/bikes", handlers.UpdateBike)

	log.Fatal(app.Listen(":3000"))
}
