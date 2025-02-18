package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/db"
)

type (
	PaymentActions interface {
		ConfirmPayment(ctx context.Context, id uuid.UUID) error
		FailPayment(ctx context.Context, id uuid.UUID) error
	}

	PaymentService struct{}
)

func NewPaymentService() PaymentActions {
	return &PaymentService{}
}

func (ps *PaymentService) ConfirmPayment(ctx context.Context, id uuid.UUID) error {
	if err := db.Queries().ConfirmPayment(ctx, id); err != nil {
		return err
	}

	return nil
}

func (ps *PaymentService) FailPayment(ctx context.Context, id uuid.UUID) error {
	if err := db.Queries().FailPayment(ctx, id); err != nil {
		return err
	}

	return nil
}
