package controller

import (
	"strconv"

	"PayWatcher/domain"

	"github.com/gofiber/fiber/v2"
)

type categoryCtrl struct {
	categoryRepo domain.CategoryRepository
}

func NewCategoryCtrl(categoryRepo domain.CategoryRepository) *categoryCtrl {
	return &categoryCtrl{
		categoryRepo: categoryRepo,
	}
}

func (ctrl categoryCtrl) GetAllCatories(c *fiber.Ctx, userID int) error {
	categories, err := ctrl.categoryRepo.GetAll(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Error al obtener categorias"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": categories, "message": "Categorias encontradas"})
}

func (ctrl categoryCtrl) GetCategoryByID(c *fiber.Ctx, ID string, userID int) error {
	iID, err := strconv.Atoi(ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"status":  "Error",
				"message": "Categoria ID no es un numero",
			},
		)
	}

	category, err := ctrl.categoryRepo.GetByID(c.Context(), iID, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Categoria no encontrada"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": category, "message": "Categorias encontradas"})
}
