package router

import (
	"PayWatcher/controller"
	"PayWatcher/database"
	"PayWatcher/middleware"
	"PayWatcher/repository/category"
	"PayWatcher/repository/payment"
	"PayWatcher/repository/user"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {

	userRepository := user.New(database.DB)
	userCtrl := controller.NewUserCtrl(userRepository)

	paymentRepository := payment.New(database.DB)
	paymentCtrl := controller.NewPaymentCtrl(paymentRepository)

	categoryRepository := category.New(database.DB)
	categoryCtrl := controller.NewCategoryCtrl(categoryRepository)

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)

	user := api.Group("/user")
	user.Post("/", userCtrl.CreateUser)
	user.Get("/", middleware.ProtectedHandler(), userCtrl.GetAllUsers)
	user.Get("/:id", middleware.ProtectedHandler(), userCtrl.GetUserByID)
	user.Put("/:id", middleware.ProtectedHandler(), userCtrl.UpdateUser)
	user.Delete("/:id", middleware.ProtectedHandler(), userCtrl.DeleteUser)

	category := api.Group("/category")
	category.Post("/", middleware.ProtectedHandler(), categoryCtrl.CreateCategory)
	category.Get("/", middleware.ProtectedHandler(), categoryCtrl.GetAllCatories)
	category.Get("/:id", middleware.ProtectedHandler(), categoryCtrl.GetCategoryByID)
	category.Put("/:id", middleware.ProtectedHandler(), categoryCtrl.UpdateCategory)
	category.Delete("/:id", middleware.ProtectedHandler(), categoryCtrl.DeleteCategory)

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
