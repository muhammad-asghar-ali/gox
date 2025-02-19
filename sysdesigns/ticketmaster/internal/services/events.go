package services

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/config"
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
		SearchEvents(ctx context.Context, req types.SearchEvent) ([]entities.Event, error)
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

	ts, err := conv.ByteToTickets(event.Tickets)
	if err != nil {
		return nil, err
	}

	res := &types.GetEventByID{
		Event:      event.Event,
		Venue:      event.Venue,
		Ticket:     ts,
		Performers: ps,
	}

	return res, nil
}

func (es *EventService) SearchEvents(ctx context.Context, req types.SearchEvent) ([]entities.Event, error) {
	args := conv.ConvertEventSearchParams(req)

	key := types.GenerateCacheKeyForSearchEvent()
	rdb := config.GetRedisClient()

	data, err := rdb.Get(ctx, key).Result()
	if err == nil && data != "" {
		cached := make([]entities.Event, 0)

		err := json.Unmarshal([]byte(data), &cached)
		if err == nil && len(cached) > 0 {
			return cached, nil
		}
	}

	log.Println("cache miss: query from database")
	events, err := db.Queries().SearchEvents(ctx, args)
	if err != nil {
		return nil, err
	}

	json, err := json.Marshal(events)
	if err != nil {
		return nil, err
	}

	err = rdb.Set(ctx, key, json, 1000*time.Microsecond).Err()
	if err != nil {
		return nil, err
	}

	return events, nil
}
