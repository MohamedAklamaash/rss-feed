-- name: CreateFeed :one
INSERT INTO Feeds (id, createdAt, updatedAt, name, url, user_id, FeedQuantity, processed)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING *;

-- name: GetUserFeeds :many
SELECT * FROM feeds WHERE user_id=$1 ORDER BY createdAt desc ;

-- name: GetAllFeeds :many
SELECT * FROM Feeds ORDER BY createdAt desc ;

-- name: GetSpecificFeed :one
SELECT * FROM Feeds where id=$1;

-- name: MarkFeedProcessed :one
UPDATE feeds
SET processed = true,
    updatedAt = NOW()
WHERE id = $1
    RETURNING *;

-- name: FeedsWithoutProcess :many
SELECT * FROM Feeds where processed=False;