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

	app.Use(middleware.PanicRecover, middleware.SecureHeaders)

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
	category.Use(middleware.ProtectedHandler())
	category.Post("/", categoryCtrl.CreateCategory)
	category.Get("/", categoryCtrl.GetAllCatories)
	category.Get("/:id", categoryCtrl.GetCategoryByID)
	category.Put("/:id", categoryCtrl.UpdateCategory)
	category.Delete("/:id", categoryCtrl.DeleteCategory)

	payment := api.Group("/payment")
	payment.Use(middleware.ProtectedHandler())
	payment.Post("/", paymentCtrl.CreatePayment)
	payment.Get("/", paymentCtrl.GetAllPayments)
	payment.Get("/:id", paymentCtrl.GetPaymentByID)
	payment.Get("/category/:idCategory", paymentCtrl.GetPaymentsByCategoryID)
	payment.Put("/:id", paymentCtrl.UpdatePayment)
	payment.Delete("/:id", paymentCtrl.DeletePayment)

	mail := api.Group("/mail")
	mail.Get("/test", middleware.ProtectedHandler(), controller.TestMail)
}
