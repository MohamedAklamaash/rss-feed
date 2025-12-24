-- name: CreateFeedFollow :one

INSERT INTO FeedFollows (id, createdAt, updatedAt, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: