-- name: AddPerformer :one
INSERT INTO performers (name, genre, bio)
VALUES ($1, $2, $3)
RETURNING *;