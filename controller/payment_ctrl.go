package controller

import (
	"PayWatcher/domain"
	"PayWatcher/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type paymentCtrl struct {
	paymentRepo domain.PaymentRepository
}

func NewPaymentCtrl(PaymentRepository domain.PaymentRepository) *paymentCtrl {
	return &paymentCtrl{
		paymentRepo: PaymentRepository,
	}
}

func (ctrl paymentCtrl) CreatePayment(c *fiber.Ctx) error {
	userID := getIdUserInToken(c)
	var params model.UpdateOrCreatePayment
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": "Error al parsear los datos"})
	}

	payment, err := ctrl.paymentRepo.Create(c.Context(), userID, params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": "Error al crear el pago"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "Success", "data": payment, "message": "Pago creado"})
}

func (ctrl paymentCtrl) DeletePayment(c *fiber.Ctx) error {
	userID := getIdUserInToken(c)
	ID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": "Error al parsear el ID a numero"})
	}

	payment, err := ctrl.paymentRepo.Delete(c.Context(), ID, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Pago no encontrado"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": payment, "message": "Pago eliminado"})
}

func (ctrl paymentCtrl) GetAllPayments(c *fiber.Ctx) error {
	userID := getIdUserInToken(c)
	payments, err := ctrl.paymentRepo.GetAll(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Error al obtener pagos"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": payments, "message": "Pagos encontrados"})
}

func (ctrl paymentCtrl) GetPaymentsByCategoryID(c *fiber.Ctx) error {
	userID := getIdUserInToken(c)
	categoryID, err := strconv.Atoi(c.Params("idCategory"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Error al parsear el ID a numero"})
	}

	payments, err := ctrl.paymentRepo.GetByCategoryID(c.Context(), categoryID, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Error al obtener pagos"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": payments, "message": "Pagos encontrados"})
}

func (ctrl paymentCtrl) GetPaymentByID(c *fiber.Ctx) error {
	userID := getIdUserInToken(c)
	ID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": "Error al parsear el ID a numero"})
	}

	payment, err := ctrl.paymentRepo.GetByID(c.Context(), ID, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Pago no encontrado"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": payment, "message": "Pago encontrado"})
}

func (ctrl paymentCtrl) UpdatePayment(c *fiber.Ctx) error {
	userID := getIdUserInToken(c)
	ID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": "Error al parsear el ID a numero"})
	}

	var params model.UpdateOrCreatePayment
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": "Error al parsear los datos"})
	}

	payment, err := ctrl.paymentRepo.Update(c.Context(), ID, userID, params)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Pago no encontrado"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": payment, "message": "Pago actualizado"})
}
