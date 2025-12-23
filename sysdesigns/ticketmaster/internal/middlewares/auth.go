package middlewares

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/config"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/db"
)

func Auth(c fiber.Ctx) error {
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
		return []byte(config.GetJwtSecret()), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id, ok := claims["user_id"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token claims"})
		}

		user_id, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token claims"})
		}

		user, err := db.Queries().GetUserByID(context.Background(), user_id)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
		}

		c.Locals("user", user)
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}
}
