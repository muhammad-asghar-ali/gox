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
	Booking interface {
		CreateBooking(c fiber.Ctx) error
		GetUserBookings(c fiber.Ctx) error
		BookTicket(c fiber.Ctx) error
		GetBookingByID(c fiber.Ctx) error
		ConfirmBooking(c fiber.Ctx) error
		CancelBooking(c fiber.Ctx) error
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

func (bh *BookingHandler) BookTicket(c fiber.Ctx) error {
	req := entities.BookTicketParams{}
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(err.Error()))
	}

	user := c.Locals("user").(entities.User)
	req.UserID = &user.ID

	if err := bh.BookingService.BookTicket(context.Background(), req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(true, "Ticket booked successfully"))
}

func (bh *BookingHandler) GetBookingByID(c fiber.Ctx) error {
	param := c.Params("id")

	id, err := uuid.Parse(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(err.Error()))
	}

	booking, err := bh.BookingService.GetBookingByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(booking, "Booking fetched successfully"))
}

func (bh *BookingHandler) ConfirmBooking(c fiber.Ctx) error {
	param := c.Params("id")

	id, err := uuid.Parse(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(err.Error()))
	}

	if err := bh.BookingService.ConfirmBooking(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(true, "Booking confirm successfully"))
}

func (bh *BookingHandler) CancelBooking(c fiber.Ctx) error {
	param := c.Params("id")

	id, err := uuid.Parse(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(err.Error()))
	}

	if err := bh.BookingService.CancelBooking(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(true, "Booking cancel successfully"))
}
