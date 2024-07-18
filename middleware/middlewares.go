package middleware

import (
	"log"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

func PanicRecover(c *fiber.Ctx) error {
	defer func() {
		if err := recover(); err != nil {
			// Registra el error
			log.Printf("Panic: %s", err)
			// Establece el encabezado de conexión a "close"
			c.Set("Connection", "close")
			// Retorna una respuesta de error 500
			c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Internal Server Error: %s", err))
		}
	}()
	return c.Next()
}
func SecureHeaders(c *fiber.Ctx) error {
	// Establece los encabezados de seguridad
	c.Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
	c.Set("Referrer-Policy", "origin-when-cross-origin")
	c.Set("X-Content-Type-Options", "nosniff")
	c.Set("X-Frame-Options", "deny")
	c.Set("X-XSS-Protection", "0")

	return c.Next()
}
