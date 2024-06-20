package router

import (
	"PayWatcher/controller"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	api := app.Group("/api")
	user := api.Group("/user")
	user.Get("/", controller.GetUser)
	user.Get("/:id", controller.GetUser)
	user.Post("/", controller.CreateUser)
	user.Put("/:id", controller.UpdateUser)
	user.Delete("/:id", controller.DeteleUser)
}
