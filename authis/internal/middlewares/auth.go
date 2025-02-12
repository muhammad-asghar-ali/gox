package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

	"authis/internal/config"
	"authis/internal/models"
)

func Auth(c *fiber.Ctx) error {
	config := config.GetConfig()

	authorization := c.Get("Authorization")
	if authorization == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "No authorization header"})
	}

	token_string := strings.Replace(authorization, "Bearer ", "", 1)
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(config.JWTSecret), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		user_id, ok := claims["user_id"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token claims"})
		}

		user := models.User{}
		user.GetUser(uint(user_id))

		if user.ID == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "User not found"})
		}

		c.Locals("user", user)
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}
}
