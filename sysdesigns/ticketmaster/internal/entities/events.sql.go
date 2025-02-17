// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: events.sql

package entities

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO events (name, description, added_by, venue_id, event_date)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, description, added_by, venue_id, event_date, created_at
`

type CreateEventParams struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	AddedBy     *uuid.UUID       `json:"added_by"`
	VenueID     uuid.UUID        `json:"venue_id"`
	EventDate   pgtype.Timestamp `json:"event_date"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.db.QueryRow(ctx, createEvent,
		arg.Name,
		arg.Description,
		arg.AddedBy,
		arg.VenueID,
		arg.EventDate,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.AddedBy,
		&i.VenueID,
		&i.EventDate,
		&i.CreatedAt,
	)
	return i, err
}

const getEventByID = `-- name: GetEventByID :one
SELECT 
    e.id, e.name, e.description, e.added_by, e.venue_id, e.event_date, e.created_at,
    v.id, v.name, v.location, v.capacity, v.added_by, v.created_at,
    JSON_AGG(
        JSON_BUILD_OBJECT(
            'id', p.id,
            'name', p.name,
            'genre', p.genre,
            'bio', p.bio
        )
    ) AS performers,
    JSON_AGG(
        JSON_BUILD_OBJECT(
            'id', t.id,
            'ticket_type', t.ticket_type,
            'price', t.price,
            'total_tickets', t.total_tickets,
            'available_tickets', t.available_tickets
        )
    ) AS tickets
FROM events e
INNER JOIN venues v ON v.id = e.venue_id
INNER JOIN tickets t ON t.event_id = e.id
INNER JOIN event_performers ep ON ep.event_id = e.id
INNER JOIN performers p ON p.id = ep.performer_id
WHERE e.id = $1
GROUP BY e.id, t.id, v.id
`

type GetEventByIDRow struct {
	Event      Event  `json:"event"`
	Venue      Venue  `json:"venue"`
	Performers []byte `json:"performers"`
	Tickets    []byte `json:"tickets"`
}

func (q *Queries) GetEventByID(ctx context.Context, id uuid.UUID) (GetEventByIDRow, error) {
	row := q.db.QueryRow(ctx, getEventByID, id)
	var i GetEventByIDRow
	err := row.Scan(
		&i.Event.ID,
		&i.Event.Name,
		&i.Event.Description,
		&i.Event.AddedBy,
		&i.Event.VenueID,
		&i.Event.EventDate,
		&i.Event.CreatedAt,
		&i.Venue.ID,
		&i.Venue.Name,
		&i.Venue.Location,
		&i.Venue.Capacity,
		&i.Venue.AddedBy,
		&i.Venue.CreatedAt,
		&i.Performers,
		&i.Tickets,
	)
	return i, err
}

const listEvent = `-- name: ListEvent :many
SELECT id, name, description, added_by, venue_id, event_date, created_at FROM events
`

func (q *Queries) ListEvent(ctx context.Context) ([]Event, error) {
	rows, err := q.db.Query(ctx, listEvent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.AddedBy,
			&i.VenueID,
			&i.EventDate,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchEvents = `-- name: SearchEvents :many
SELECT id, name, description, added_by, venue_id, event_date, created_at
FROM events
WHERE (name ILIKE '%' || $1 || '%' OR description ILIKE '%' || $1 || '%')
  AND event_date BETWEEN $2 AND $3
ORDER BY event_date ASC
LIMIT $4
OFFSET $5
`

type SearchEventsParams struct {
	Column1     pgtype.Text      `json:"column_1"`
	EventDate   pgtype.Timestamp `json:"event_date"`
	EventDate_2 pgtype.Timestamp `json:"event_date_2"`
	Limit       int32            `json:"limit"`
	Offset      int32            `json:"offset"`
}

func (q *Queries) SearchEvents(ctx context.Context, arg SearchEventsParams) ([]Event, error) {
	rows, err := q.db.Query(ctx, searchEvents,
		arg.Column1,
		arg.EventDate,
		arg.EventDate_2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.AddedBy,
			&i.VenueID,
			&i.EventDate,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
