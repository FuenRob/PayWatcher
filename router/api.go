package router

import (
	"PayWatcher/controller"
	"PayWatcher/database"
	"PayWatcher/middleware"
	"PayWatcher/repository/payment"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {

	paymentRepository := payment.New(database.DB)
	paymentCtrl := controller.NewPaymentCtrl(paymentRepository)

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

	payment := api.Group("/payment")
	payment.Post("/", middleware.ProtectedHandler(), paymentCtrl.CreatePayment)
	payment.Get("/", middleware.ProtectedHandler(), paymentCtrl.GetAllPayments)
	payment.Get("/:id", middleware.ProtectedHandler(), paymentCtrl.GetPaymentByID)
	payment.Get("/category/:idCategory", middleware.ProtectedHandler(), paymentCtrl.GetPaymentsByCategoryID)
	payment.Put("/:id", middleware.ProtectedHandler(), paymentCtrl.UpdatePayment)
	payment.Delete("/:id", middleware.ProtectedHandler(), paymentCtrl.DeletePayment)

	mail := api.Group("/mail")
	mail.Get("/test", middleware.ProtectedHandler(), controller.TestMail)
}
