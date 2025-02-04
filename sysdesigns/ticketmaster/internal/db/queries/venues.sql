-- name: CreateVenue :one
INSERT INTO venues (name, location, capacity, added_by)
VALUES ($1, $2, $3, $4)
RETURNING *;
