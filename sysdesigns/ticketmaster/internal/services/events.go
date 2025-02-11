package services

import (
	"context"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/db"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

type (
	EventActions interface {
		CreateEvent(ctx context.Context, req entities.CreateEventParams) (*entities.Event, error)
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
