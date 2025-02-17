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
	Ticket interface {
		CreateTicket(c fiber.Ctx) error
		GetAvailableTickets(c fiber.Ctx) error
	}

	TicketHandler struct {
		TicketService services.TicketService
	}
)

func NewTicketHandler(ts services.TicketService) Ticket {
	return &TicketHandler{TicketService: ts}
}

func (th *TicketHandler) CreateTicket(c fiber.Ctx) error {
	req := entities.CreateTicketParams{}
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(err.Error()))
	}

	created, err := th.TicketService.CreateTicket(context.Background(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse(created, "Ticket created successfully"))
}

func (th *TicketHandler) GetAvailableTickets(c fiber.Ctx) error {
	param := c.Params("id")

	id, err := uuid.Parse(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(err.Error()))
	}

	count, err := th.TicketService.GetAvailableTickets(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(count, "Available tickets successfully"))
}
