package controller

import (
	"strconv"

	"PayWatcher/domain"
	"PayWatcher/model"

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

func (ctrl categoryCtrl) GetAllCatories(c *fiber.Ctx) error {
	userID := getIdUserInToken(c)
	categories, err := ctrl.categoryRepo.GetAll(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Error al obtener categorias"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": categories, "message": "Categorias encontradas"})
}

func (ctrl categoryCtrl) GetCategoryByID(c *fiber.Ctx) error {
	userID := getIdUserInToken(c)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"status":  "Error",
				"message": "Categoria ID no es un numero",
			},
		)
	}

	category, err := ctrl.categoryRepo.GetByID(c.Context(), id, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Categoria no encontrada"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": category, "message": "Categorias encontradas"})
}

func (ctrl categoryCtrl) CreateCategory(c *fiber.Ctx) error {
	userID := getIdUserInToken(c)
	params := new(model.UpdateOrCreateCategory)
	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": "Error al parsear los datos"})
	}

	category, err := ctrl.categoryRepo.Create(c.Context(), userID, *params)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Error al crear categoria"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": category, "message": "Categoria creada"})
}

func (ctrl categoryCtrl) UpdateCategory(c *fiber.Ctx) error {
	userID := getIdUserInToken(c)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"status":  "Error",
				"message": "Categoria ID no es un numero",
			},
		)
	}

	params := new(model.UpdateOrCreateCategory)
	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": "Error al parsear los datos"})
	}

	category, err := ctrl.categoryRepo.Update(c.Context(), id, userID, *params)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Error al actualizar categoria"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": category, "message": "Categoria actualizada"})
}

func (ctrl categoryCtrl) DeleteCategory(c *fiber.Ctx) error {
	userID := getIdUserInToken(c)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"status":  "Error",
				"message": "Categoria ID no es un numero",
			},
		)
	}

	category, err := ctrl.categoryRepo.Delete(c.Context(), id, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Error al eliminar categoria"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": category, "message": "Categoria eliminada"})
}
