package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/services"
)

type (
	Event interface {
		CreateEvent(c fiber.Ctx) error
		ListEvent(c fiber.Ctx) error
		GetEventByID(c fiber.Ctx) error
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

func (eh *EventHandler) ListEvent(c fiber.Ctx) error {
	list, err := eh.EventService.ListEvent(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(list, "Events fetched successfully"))
}

func (eh *EventHandler) GetEventByID(c fiber.Ctx) error {
	id := c.Params("id")

	user_id, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(err.Error()))
	}

	event, err := eh.EventService.GetEventByID(context.Background(), user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(event, "Event fetched successfully"))
}
