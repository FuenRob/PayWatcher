package controller

import (
	"PayWatcher/database"
	"PayWatcher/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GET /users
// GET /user/:id
func GetUser(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")

	if id != "" {
		return getUserByID(c, db, id)
	}

	return getAllUsers(c, db)
}

func getUserByID(c *fiber.Ctx, db *gorm.DB, id string) error {
	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Usuario no encontrado"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": user, "message": "Usuario encontrado"})
}

func getAllUsers(c *fiber.Ctx, db *gorm.DB) error {
	var users []model.User
	if err := db.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Error", "message": "Error al obtener usuarios"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": users, "message": "Usuarios encontrados"})
}

// POST /user
func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": "Faltan datos en el JSON"})
	}

	password, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Error", "message": "Error al cifrar la contraseña"})
	}

	user.Password = password

	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Error", "message": "Error al crear el usuario"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": user, "message": "Usuario creado correctamente"})
}

// PUT /user/:id
func UpdateUser(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var user model.User

	type updateUser struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	var uu updateUser
	if err := c.BodyParser(&uu); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": "Faltan datos en el JSON"})
	}

	password, err := hashPassword(uu.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Error", "message": "Error al cifrar la contraseña"})
	}

	if err := db.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Usuario no encontrado"})
	}

	user.Name = uu.Name
	user.Email = uu.Email
	user.UserName = uu.UserName
	user.Password = password

	if err := db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Error", "message": "Error al actualizar el usuario"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": user, "message": "Usuario editado correctamente"})
}

// DELETE /user/:id
func DeteleUser(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var user model.User

	if err := db.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "Usuario no encontrado"})
	}

	if err := db.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Error", "message": "Error al borrar el usuario"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "data": user, "message": "Usuario borrado correctamente"})
}
