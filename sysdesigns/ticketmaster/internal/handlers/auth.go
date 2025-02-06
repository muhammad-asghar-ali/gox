package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/services"
)

type (
	Auth interface {
		Register(c fiber.Ctx) error
		Login(c fiber.Ctx) error
	}

	AuthHandler struct {
		UserService services.UserService
	}
)

func NewAuthHandler(us services.UserService) Auth {
	return &AuthHandler{UserService: us}
}

func (uh *AuthHandler) Register(c fiber.Ctx) error {
	req := entities.CreateUserParams{}
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	user, err := uh.UserService.Create(context.Background(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "User register successfully",
		"data":    user,
	})
}

func (uh *AuthHandler) Login(c fiber.Ctx) error {
	req := common.LoginRequest{}
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	user, err := uh.UserService.Login(context.Background(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "User login successfully",
		"data":    user,
	})
}
