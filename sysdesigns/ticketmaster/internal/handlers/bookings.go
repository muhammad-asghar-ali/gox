package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/services"
)

type (
	Booking interface {
		CreateBooking(c fiber.Ctx) error
		GetUserBookings(c fiber.Ctx) error
	}

	BookingHandler struct {
		BookingService services.BookingService
	}
)

func NewBookingHandler(bs services.BookingService) Booking {
	return &BookingHandler{BookingService: bs}
}

func (bh *BookingHandler) CreateBooking(c fiber.Ctx) error {
	req := entities.CreateBookingParams{}
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(err.Error()))
	}

	user := c.Locals("user").(entities.User)
	req.UserID = &user.ID

	created, err := bh.BookingService.CreateBooking(context.Background(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse(created, "Booking created successfully"))
}

func (bh *BookingHandler) GetUserBookings(c fiber.Ctx) error {
	user := c.Locals("user").(entities.User)

	bookings, err := bh.BookingService.GetUserBookings(context.Background(), &user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(bookings, "User Bookings fetched successfully"))
}
