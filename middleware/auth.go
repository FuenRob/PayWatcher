package middleware

import (
	"PayWatcher/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func ProtectedHandler() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.SecretJWTKey)},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Error", "message": "Missing or malformed JWT",
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status": "Error", "message": "Invalid or expired JWT",
	})
}
