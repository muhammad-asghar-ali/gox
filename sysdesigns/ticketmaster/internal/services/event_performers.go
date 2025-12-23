package services

import (
	"context"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/db"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

type (
	EventPerformerActions interface {
		AddPerformerToEvent(ctx context.Context, req entities.AddPerformerToEventParams) (*entities.EventPerformer, error)
	}

	EventPerformerService struct{}
)

func NewEventPerformerService() EventPerformerActions {
	return &EventPerformerService{}
}

func (eps *EventPerformerService) AddPerformerToEvent(ctx context.Context, req entities.AddPerformerToEventParams) (*entities.EventPerformer, error) {
	ep, err := db.Queries().AddPerformerToEvent(ctx, req)
	if err != nil {
		return nil, err
	}

	return &ep, nil
}
