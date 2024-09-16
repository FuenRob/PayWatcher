package controller

import (
	"PayWatcher/config"

	"github.com/gofiber/fiber/v2"
)

func TestMail(c *fiber.Ctx) error {
	error := config.MailSender.SendMail([]string{"fuenrob@gmail.com"}, "Prueba de envio", "El cuerpo del correo")
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Error", "message": "Error al enviar el email."})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "message": "Email enviado correctamente."})
}
