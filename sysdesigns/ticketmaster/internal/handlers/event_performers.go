package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/services"
)

type (
	EventPerformer interface {
		AddPerformerToEvent(c fiber.Ctx) error
	}

	EventPerformerHandler struct {
		EventPerformerService services.EventPerformerService
	}
)

func NewEventPerformerHandler(eps services.EventPerformerService) EventPerformer {
	return &EventPerformerHandler{EventPerformerService: eps}
}

func (eph *EventPerformerHandler) AddPerformerToEvent(c fiber.Ctx) error {
	req := entities.AddPerformerToEventParams{}
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(err.Error()))
	}

	added, err := eph.EventPerformerService.AddPerformerToEvent(context.Background(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse(added, "Performer added to event successfully"))
}
