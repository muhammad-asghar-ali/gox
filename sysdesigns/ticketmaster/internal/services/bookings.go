package services

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common/enums"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common/types"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/db"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

type (
	BookingActions interface {
		CreateBooking(ctx context.Context, req entities.CreateBookingParams) (*entities.Booking, error)
		GetUserBookings(ctx context.Context, userID *uuid.UUID) ([]entities.Booking, error)
		BookTicket(ctx context.Context, req entities.BookTicketParams) error
		GetBookingByID(ctx context.Context, id uuid.UUID) (*types.GetBookingByID, error)
		ConfirmBooking(ctx context.Context, id uuid.UUID) error
		CancelBooking(ctx context.Context, id uuid.UUID) error
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

func (bs *BookingService) GetUserBookings(ctx context.Context, userID *uuid.UUID) ([]entities.Booking, error) {
	bookings, err := db.Queries().GetUserBookings(ctx, userID)
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (bs *BookingService) BookTicket(ctx context.Context, req entities.BookTicketParams) error {
	txConn := db.BeginTx(ctx)
	tx := db.Queries().WithTx(txConn)

	check, err := tx.GetRemainingTickets(ctx, req.TicketID)
	if err != nil {
		return nil
	}

	if check.EventAvailable >= req.Quantity && check.TicketAvailable >= req.Quantity {
		if err := tx.BookTicket(ctx, req); err != nil {
			return err
		}

		booking, err := tx.GetLastBooking(ctx, req.UserID)
		if err != nil {
			return err
		}

		preq := entities.CreatePaymentParams{
			BookingID:     booking.ID,
			UserID:        booking.UserID,
			Amount:        req.TotalPrice,
			PaymentMethod: enums.PaymentMethodCreditCard.String(),
			Status:        enums.PaymentStatusPending.String(),
		}

		// create payment
		_, err = tx.CreatePayment(ctx, preq)
		if err != nil {
			return err
		}

		if err := txConn.Commit(ctx); err != nil {
			return err
		}

		return nil
	}

	txConn.Rollback(ctx)
	return errors.New("tickets are not available")
}

func (bs *BookingService) GetBookingByID(ctx context.Context, id uuid.UUID) (*types.GetBookingByID, error) {
	row, err := db.Queries().GetBookingByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := &types.GetBookingByID{
		Booking: row.Booking,
		Ticket:  row.Ticket,
		Payment: row.Payment,
	}

	return res, nil
}

func (bs *BookingService) ConfirmBooking(ctx context.Context, id uuid.UUID) error {
	status, err := db.Queries().GetBookingStatus(ctx, id)
	if err != nil {
		return err
	}

	if status != enums.BookingStatusPending.String() {
		return errors.New("unable to update booking status")
	}

	if err := db.Queries().ConfirmBooking(ctx, id); err != nil {
		return err
	}

	return nil
}

func (bs *BookingService) CancelBooking(ctx context.Context, id uuid.UUID) error {
	status, err := db.Queries().GetBookingStatus(ctx, id)
	if err != nil {
		return err
	}

	if status != enums.BookingStatusPending.String() {
		return errors.New("unable to update booking status")
	}

	if err := db.Queries().CancelBooking(ctx, id); err != nil {
		return err
	}

	return nil
}
