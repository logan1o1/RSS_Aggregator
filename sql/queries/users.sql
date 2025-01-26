-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_it, name)
VALUES($1, $2, $3, $4)
RETURNING *;