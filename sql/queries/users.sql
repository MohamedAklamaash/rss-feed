-- name: CreateUser :one
INSERT INTO users (id, createdAt, updatedAt, name)
VALUES ($1, $2, $3, $4)
    RETURNING *;
