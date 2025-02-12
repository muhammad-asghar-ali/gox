package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/services"
)

type (
	Event interface {
		CreateEvent(c fiber.Ctx) error
	}

	EventHandler struct {
		EventService services.EventService
	}
)

func NewEventHandler(es services.EventService) Event {
	return &EventHandler{EventService: es}
}

func (eh *EventHandler) CreateEvent(c fiber.Ctx) error {
	req := entities.CreateEventParams{}
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(err.Error()))
	}

	user := c.Locals("user").(entities.User)
	req.AddedBy = &user.ID

	created, err := eh.EventService.CreateEvent(context.Background(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse(created, "Event add successfully"))
}
