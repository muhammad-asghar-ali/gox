package services

import (
	"context"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/db"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

type (
	VenueActions interface {
		CreateVenue(ctx context.Context, req entities.CreateVenueParams) (*entities.Venue, error)
	}

	VenueService struct{}
)

func NewVenueService() VenueActions {
	return &VenueService{}
}

func (vs *VenueService) CreateVenue(ctx context.Context, req entities.CreateVenueParams) (*entities.Venue, error) {
	venue, err := db.Queries().CreateVenue(ctx, req)
	if err != nil {
		return nil, err
	}

	return &venue, nil
}
