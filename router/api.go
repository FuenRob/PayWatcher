package router

import (
	"PayWatcher/controller"
	"PayWatcher/middleware"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	//add panic recover
	app.Use(middleware.PanicRecover, middleware.SecureHeaders)

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
	category.Use(middleware.ProtectedHandler())
	category.Post("/", controller.CreateCategory)
	category.Get("/", controller.GetCatories)
	category.Get("/:id", controller.GetCatories)
	category.Put("/:id", controller.UpdateCategory)
	category.Delete("/:id", controller.DeleteCategory)

	payment := api.Group("/payment")
	payment.Use(middleware.ProtectedHandler())
	payment.Post("/", controller.CreatePayment)
	payment.Get("/", controller.GetPayment)
	payment.Get("/:id", controller.GetPayment)
	payment.Get("/category/:idCategory", controller.GetPaymentsByCategoryID)
	payment.Put("/:id", controller.UpdatePayment)
	payment.Delete("/:id", controller.DeletePayment)
}
