// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: venues.sql

package entities

import (
	"context"

	"github.com/google/uuid"
)

const createVenue = `-- name: CreateVenue :one
INSERT INTO venues (name, location, capacity, added_by)
VALUES ($1, $2, $3, $4)
RETURNING id, name, location, capacity, added_by, created_at
`

type CreateVenueParams struct {
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Capacity int32     `json:"capacity"`
	AddedBy *uuid.UUID `json:"added_by"`
}

// @optional added_by
func (q *Queries) CreateVenue(ctx context.Context, arg CreateVenueParams) (Venue, error) {
	row := q.db.QueryRow(ctx, createVenue,
		arg.Name,
		arg.Location,
		arg.Capacity,
		arg.AddedBy,
	)
	var i Venue
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.Capacity,
		&i.AddedBy,
		&i.CreatedAt,
	)
	return i, err
}

const listVenue = `-- name: ListVenue :many
SELECT id, name, location, capacity, added_by, created_at FROM venues
`

func (q *Queries) ListVenue(ctx context.Context) ([]Venue, error) {
	rows, err := q.db.Query(ctx, listVenue)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Venue
	for rows.Next() {
		var i Venue
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Location,
			&i.Capacity,
			&i.AddedBy,
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
