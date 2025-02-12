package services

import (
	"context"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/db"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

type (
	BookingActions interface {
		CreateBooking(ctx context.Context, req entities.CreateBookingParams) (*entities.Booking, error)
	}

	BookingService struct{}
)

func NewBookingService() BookingActions {
	return &BookingService{}
}

func (bs *BookingService) CreateBooking(ctx context.Context, req entities.CreateBookingParams) (*entities.Booking, error) {
	b, err := db.Queries().CreateBooking(ctx, req)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
