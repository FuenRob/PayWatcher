package middleware

import (
	"log"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

func PanicRecover(c *fiber.Ctx) error {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic: %s", err)
			c.Set("Connection", "close")
			c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Internal Server Error: %s", err))
		}
	}()
	return c.Next()
}
func SecureHeaders(c *fiber.Ctx) error {
	c.Set("Referrer-Policy", "origin-when-cross-origin")
	c.Set("X-Content-Type-Options", "nosniff")

	return c.Next()
}
