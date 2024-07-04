package router

import (
	"PayWatcher/controller"
	"PayWatcher/middleware"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)

	user := api.Group("/user")
	user.Post("/", controller.CreateUser)
	user.Get("/", middleware.ProtectedHandler(), controller.GetUser)
	user.Get("/:id", middleware.ProtectedHandler(), controller.GetUser)
	user.Put("/:id", middleware.ProtectedHandler(), controller.UpdateUser)
	user.Delete("/:id", middleware.ProtectedHandler(), controller.DeteleUser)

	category := api.Group("/category")
	category.Post("/", middleware.ProtectedHandler(), controller.CreateCategory)
	category.Get("/", middleware.ProtectedHandler(), controller.GetCatories)
	category.Get("/:id", middleware.ProtectedHandler(), controller.GetCatories)
	category.Put("/:id", middleware.ProtectedHandler(), controller.UpdateCategory)
	category.Delete("/:id", middleware.ProtectedHandler(), controller.DeleteCategory)
}
