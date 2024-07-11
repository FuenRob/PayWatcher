package controller

import (
	"PayWatcher/config"
	"PayWatcher/database"
	"PayWatcher/model"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GET /payments
// GET /payments/:id
func GetPayment(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	userID := getIdUserInToken(c)

	if id != "" {
		return getPaymentByID(c, db, id, userID)
	}

	return getAllPayments(c, db, userID)
}

func getAllPayments(c *fiber.Ctx, db *gorm.DB, userID int) error {
	payments := []model.Payment{}
	if err := db.Where("user_id = ?", userID).Find(&payments).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "data": err, "message": "Error al obtener los pagos"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": payments, "message": "Pagos encontrados correctamente"})
}

func getPaymentByID(c *fiber.Ctx, db *gorm.DB, id string, userID int) error {
	payment := model.Payment{}

	if err := db.First(&payment, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "data": err, "message": "Pago no encontrado"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": payment, "message": "Pago encontrado correctamente"})
}

// GET /payments/category/:idCategory
func GetPaymentsByCategoryID(c *fiber.Ctx) error {
	db := database.DB
	idCategory := c.Params("idCategory")
	userID := getIdUserInToken(c)
	payments := []model.Payment{}
	if err := db.Where("user_id = ? AND category_id = ?", userID, idCategory).Find(&payments).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "data": err, "message": fmt.Sprintf("Error al obtener los pagos de la categoria %s", idCategory)})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": payments, "message": "Pagos encontrados correctamente"})
}

// POST /payments
func CreatePayment(c *fiber.Ctx) error {
	db := database.DB
	payment := new(model.Payment)
	createPaymentStruct := new(model.UpdateOrCreatePayment)
	if err := c.BodyParser(&createPaymentStruct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "data": err, "message": "Faltan datos en el JSON"})
	}

	date, err := time.Parse(config.DateFormat, createPaymentStruct.ChargeDate)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "data": err, "message": "La fecha está en un formato erroneo"})
	}

	payment.UserID = uint(getIdUserInToken(c))
	payment.Name = createPaymentStruct.Name
	payment.CategoryID = createPaymentStruct.CategoryID
	payment.NetAmount = createPaymentStruct.NetAmount
	payment.GrossAmount = createPaymentStruct.GrossAmount
	payment.Deductible = createPaymentStruct.Deductible
	payment.ChargeDate = date
	payment.Recurrent = createPaymentStruct.Recurrent
	payment.PaymentType = createPaymentStruct.PaymentType
	payment.Paid = createPaymentStruct.Paid

	if err := db.Create(&payment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Error", "data": err, "message": "Error al crear el pago"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": payment, "message": "Pago creado correctamente"})
}

// PUT /payments/:id
func UpdatePayment(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var payment model.Payment
	var updatePaymentStruct model.UpdateOrCreatePayment

	if err := c.BodyParser(&updatePaymentStruct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "data": err, "message": "Faltan datos en el JSON"})
	}

	date, err := time.Parse(config.DateFormat, updatePaymentStruct.ChargeDate)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "data": err, "message": "La fecha está en un formato erroneo"})
	}

	if err := db.First(&payment, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "data": err, "message": "Pago no encontrado"})
	}

	payment.Name = updatePaymentStruct.Name
	payment.CategoryID = updatePaymentStruct.CategoryID
	payment.NetAmount = updatePaymentStruct.NetAmount
	payment.GrossAmount = updatePaymentStruct.GrossAmount
	payment.Deductible = updatePaymentStruct.Deductible
	payment.ChargeDate = date
	payment.Recurrent = updatePaymentStruct.Recurrent
	payment.PaymentType = updatePaymentStruct.PaymentType
	payment.Paid = updatePaymentStruct.Paid

	if err := db.Save(&payment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Error", "data": err, "message": "Error al actualizar el pago"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": payment, "message": "Pago editada correctamente"})
}

// DELETE /payments/:id
func DeletePayment(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var payment model.Payment

	if err := db.First(&payment, "id = ? AND user_id = ?", id, getIdUserInToken(c)).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "data": err, "message": "Pago no encontrado"})
	}

	if err := db.Delete(&payment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Error", "data": err, "message": "Erro al borrar el pago"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "message": "Pago eliminado correctamente"})
}
