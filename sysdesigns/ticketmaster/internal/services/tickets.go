package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/db"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

type (
	TicketActions interface {
		CreateTicket(ctx context.Context, req entities.CreateTicketParams) (*entities.Ticket, error)
		GetAvailableTickets(ctx context.Context, id uuid.UUID) (int32, error)
	}

	TicketService struct{}
)

func NewTicketService() TicketActions {
	return &TicketService{}
}

func (ts *TicketService) CreateTicket(ctx context.Context, req entities.CreateTicketParams) (*entities.Ticket, error) {
	t, err := db.Queries().CreateTicket(ctx, req)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (ts *TicketService) GetAvailableTickets(ctx context.Context, id uuid.UUID) (int32, error) {
	count, err := db.Queries().GetAvailableTickets(ctx, id)
	if err != nil {
		return 0, err
	}

	return count, nil
}
