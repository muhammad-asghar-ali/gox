package handlers

import (
	"github.com/gofiber/fiber/v3"

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
	return nil
}

func (uh *AuthHandler) Login(c fiber.Ctx) error {
	return nil
}
