-- name: CreateVenue :one
-- @optional added_by
INSERT INTO venues (name, location, capacity, added_by)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListVenue :many
SELECT * FROM venues;