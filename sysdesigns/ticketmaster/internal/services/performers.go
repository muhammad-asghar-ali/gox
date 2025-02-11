package services

import (
	"context"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/db"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

type (
	PerformerActions interface {
		AddPerformer(ctx context.Context, req entities.AddPerformerParams) (*entities.Performer, error)
	}

	PerformerService struct{}
)

func NewPerformerService() PerformerActions {
	return &PerformerService{}
}

func (ps *PerformerService) AddPerformer(ctx context.Context, req entities.AddPerformerParams) (*entities.Performer, error) {
	e, err := db.Queries().AddPerformer(ctx, req)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
