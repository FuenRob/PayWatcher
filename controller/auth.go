package controller

import (
	"PayWatcher/config"
	"PayWatcher/database"
	"PayWatcher/model"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	passCifrate, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(passCifrate), err
}

func getIdUserInToken(c *fiber.Ctx) int {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	return int(id)
}

func createToken(user model.User) *jwt.Token {
	claims := jwt.MapClaims{
		"name":  user.Name,
		"id":    user.ID,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func CheckExistingUser(identity string) (model.User, error) {
	var DB = database.DB
	var user model.User
	if err := DB.Where("user_name = ?", identity).First(&user); err.Error != nil {
		user.ID = 0
		return user, err.Error
	}

	return user, nil
}

func CheckComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// /api/auth/login post
func Login(c *fiber.Ctx) error {
	var loginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}

	var user model.User
	var err error

	if err = c.BodyParser(&loginInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid Request",
		})
	}

	if strings.Contains(loginInput.Identity, "@") {
		user, err = CheckExistingUserByEmail(loginInput.Identity)
		if err != nil || user.ID == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid credentials",
			})
		}

	} else {

		user, err = CheckExistingUser(loginInput.Identity)
		if err != nil || user.ID == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid credentials",
			})
		}
	}

	if !CheckComparePassword(user.Password, loginInput.Password) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid credentials",
		})
	}

	token := createToken(user)

	t, err := token.SignedString([]byte(config.SecretJWTKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error interno",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login OK!",
		"data":    t,
	})
}

func CheckExistingUserByEmail(email string) (model.User, error) {
	var DB = database.DB
	var user model.User
	if err := DB.Where("email = ?", email).First(&user); err.Error != nil {
		user.ID = 0
		return user, err.Error
	}

	return user, nil
}
