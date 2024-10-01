package main

import (
	"PayWatcher/cronjob"
	"PayWatcher/database"
	"PayWatcher/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	go cronjob.InitCronJobs()
	database.Connect()
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	router.Init(app)
	app.Listen(":3000")
}
