-- name: CreatePost :one
INSERT INTO Posts(id, title, description, publishedat, link, FeedId)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetAllPostsWithFeed :many
SELECT * FROM Posts WHERE FeedId = $1;