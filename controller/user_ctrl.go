package controller

import (
	"PayWatcher/domain"
	"PayWatcher/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type userCtrl struct {
	userRepo domain.UserRepository
}

func NewUserCtrl(UserRepository domain.UserRepository) *userCtrl {
	return &userCtrl{
		userRepo: UserRepository,
	}
}

func (ctrl userCtrl) GetAllUsers(c *fiber.Ctx) error {
	users, err := ctrl.userRepo.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Error al obtener usuarios"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": users, "message": "Usuarios encontrados"})
}

func (ctrl userCtrl) GetUserByID(c *fiber.Ctx) error {
	ID, _ := strconv.Atoi(c.Params("id"))
	user, err := ctrl.userRepo.GetByID(c.Context(), ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Usuario no encontrado"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": user, "message": "Usuario encontrado"})
}

func (ctrl userCtrl) UpdateUser(c *fiber.Ctx) error {
	ID, _ := strconv.Atoi(c.Params("id"))
	var params model.UpdateOrCreateUser
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": "Error al parsear los datos"})
	}

	user, err := ctrl.userRepo.Update(c.Context(), ID, params)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Usuario no encontrado"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": user, "message": "Usuario actualizado"})
}

func (ctrl userCtrl) CreateUser(c *fiber.Ctx) error {
	var params model.UpdateOrCreateUser
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": "Error al parsear los datos"})
	}

	_, err := ctrl.userRepo.Create(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Error", "message": "Error al crear el usuario"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": params, "message": "Usuario creado"})
}

func (ctrl userCtrl) DeleteUser(c *fiber.Ctx) error {
	ID, _ := strconv.Atoi(c.Params("id"))
	_, err := ctrl.userRepo.Delete(c.Context(), ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Usuario no encontrado"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "message": "Usuario eliminado"})
}
