-- name: CreateUser :one
INSERT INTO users (name, email, password, phone)
VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = $1;