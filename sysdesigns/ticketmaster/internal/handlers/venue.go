package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/services"
)

type (
	Venue interface {
		CreateVenue(c fiber.Ctx) error
	}

	VenueHandler struct {
		VenueService services.VenueService
	}
)

func NewVenueHandler(us services.VenueService) Venue {
	return &VenueHandler{VenueService: us}
}

func (vh *VenueHandler) CreateVenue(c fiber.Ctx) error {
	req := entities.CreateVenueParams{}
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(err.Error()))
	}

	user := c.Locals("user").(entities.User)
	req.AddedBy = user.ID

	created, err := vh.VenueService.CreateVenue(context.Background(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse(created, "Venue add successfully"))
}
