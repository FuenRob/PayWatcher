package controller

import (
	"PayWatcher/database"
	"PayWatcher/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GET /categories
// GET /categories/:id
func GetCatories(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	userID := getIdUserInToken(c)

	if id != "" {
		return getCategoryByID(c, db, id, userID)
	}

	return getAllCatories(c, db, userID)
}

func getAllCatories(c *fiber.Ctx, db *gorm.DB, userID int) error {
	categories := []model.Category{}
	if err := db.Where("user_id = ?", userID).Find(&categories).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Error al obtener categorias"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": categories, "message": "Categorias encontradas"})

}

func getCategoryByID(c *fiber.Ctx, db *gorm.DB, id string, userID int) error {
	var category model.Category

	if err := db.First(&category, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Categoria no encontrada"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": category, "message": "Categorias encontradas"})
}

// POST /categories
func CreateCategory(c *fiber.Ctx) error {
	db := database.DB
	category := new(model.Category)
	createCategoryStruct := new(model.UpdateOrCreateCategory)

	if err := c.BodyParser(&createCategoryStruct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": "Faltan datos en el JSON"})
	}

	category.UserID = uint(getIdUserInToken(c))
	category.Name = createCategoryStruct.Name
	category.Priority = createCategoryStruct.Priority
	category.Recurrent = createCategoryStruct.Recurrent
	category.Notify = createCategoryStruct.Notify

	if err := db.Create(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Error", "message": "Error al crear la categoria"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": category, "message": "Categoria creada correctamente"})
}

// PUT /categories/:id
func UpdateCategory(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var category model.Category

	var uc model.UpdateOrCreateCategory

	if err := c.BodyParser(&uc); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": "Faltan datos en el JSON"})
	}

	if err := db.First(&category, "id = ? AND user_id = ?", id, getIdUserInToken(c)).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Categoria no encontrada"})
	}

	category.Name = uc.Name
	category.Priority = uc.Priority
	category.Recurrent = uc.Recurrent
	category.Notify = uc.Notify

	if err := db.Save(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Error", "message": "Error al actualizar la categoria"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": category, "message": "Categoria editada correctamente"})

}

// DETELE /categories/:id
func DeleteCategory(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var category model.Category

	if err := db.First(&category, "id = ? AND user_id = ?", id, getIdUserInToken(c)).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Categoria no encontrado"})
	}

	if err := db.Delete(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Error", "message": "Error al borrar la categoria"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": category, "message": "Categoria borrada correctamente"})
}
