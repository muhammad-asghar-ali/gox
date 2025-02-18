package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/services"
)

type (
	Payment interface {
		ConfirmPayment(c fiber.Ctx) error
		FailPayment(c fiber.Ctx) error
	}

	PaymentHandler struct {
		PaymentService services.PaymentService
	}
)

func NewPaymentHandler(ps services.PaymentService) Payment {
	return &PaymentHandler{PaymentService: ps}
}

func (ph *PaymentHandler) ConfirmPayment(c fiber.Ctx) error {
	param := c.Params("id")

	id, err := uuid.Parse(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(err.Error()))
	}

	if err := ph.PaymentService.ConfirmPayment(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(true, "Payment confirmed successfully"))
}

func (ph *PaymentHandler) FailPayment(c fiber.Ctx) error {
	param := c.Params("id")

	id, err := uuid.Parse(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(err.Error()))
	}

	if err := ph.PaymentService.ConfirmPayment(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(true, "Payment failed"))
}
