package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common/conv"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common/types"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/db"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

type (
	EventActions interface {
		CreateEvent(ctx context.Context, req entities.CreateEventParams) (*entities.Event, error)
		ListEvent(ctx context.Context) ([]entities.Event, error)
		GetEventByID(ctx context.Context, id uuid.UUID) (*types.GetEventByID, error)
	}

	EventService struct{}
)

func NewEventService() EventActions {
	return &EventService{}
}

func (es *EventService) CreateEvent(ctx context.Context, req entities.CreateEventParams) (*entities.Event, error) {
	e, err := db.Queries().CreateEvent(ctx, req)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (es *EventService) ListEvent(ctx context.Context) ([]entities.Event, error) {
	events, err := db.Queries().ListEvent(ctx)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (es *EventService) GetEventByID(ctx context.Context, id uuid.UUID) (*types.GetEventByID, error) {
	event, err := db.Queries().GetEventByID(ctx, id)
	if err != nil {
		return nil, err
	}

	ps, err := conv.ByteToPerformers(event.Performers)
	if err != nil {
		return nil, err
	}

	res := &types.GetEventByID{
		Event:      event.Event,
		Ticket:     event.Ticket,
		Performers: ps,
	}

	return res, nil
}
