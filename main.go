package main

import (
	"PayWatcher/database"
	"PayWatcher/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	router.Init(app)
	app.Listen(":3000")
}
