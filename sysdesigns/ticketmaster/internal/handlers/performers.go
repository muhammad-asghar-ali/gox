package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/services"
)

type (
	Performer interface {
		AddPerformer(c fiber.Ctx) error
		ListPerformer(c fiber.Ctx) error
	}

	PerformerHandler struct {
		PerformerService services.PerformerService
	}
)

func NewPerformerHandler(ps services.PerformerService) Performer {
	return &PerformerHandler{PerformerService: ps}
}

func (ph *PerformerHandler) AddPerformer(c fiber.Ctx) error {
	req := entities.AddPerformerParams{}
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(err.Error()))
	}

	created, err := ph.PerformerService.AddPerformer(context.Background(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse(created, "Performer add successfully"))
}

func (ph *PerformerHandler) ListPerformer(c fiber.Ctx) error {
	list, err := ph.PerformerService.ListPerformer(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse(list, "Performers fetched successfully"))
}
