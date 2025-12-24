-- name: CreateUser :one
INSERT INTO users (id, createdAt, updatedAt, name, api_key)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea),'hex'))
    RETURNING *;
-- name: GetUserByAPIKey :one

SELECT * FROM Users where api_key=$1;