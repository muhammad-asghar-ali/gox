package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/db"
)

type (
	HealthHandler struct{}
)

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (hh *HealthHandler) HealthCheck(c fiber.Ctx) error {
	if err := db.Get().Ping(context.Background()); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse("Database connection failed"))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("ok", "Server and database are running"))
}
