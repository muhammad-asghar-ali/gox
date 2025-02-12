package services

import (
	"context"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/db"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

type (
	TicketActions interface {
		CreateTicket(ctx context.Context, req entities.CreateTicketParams) (*entities.Ticket, error)
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
