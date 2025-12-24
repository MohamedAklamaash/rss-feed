-- name: CreateFeedFollow :one

INSERT INTO FeedFollows (id, createdAt, updatedAt, user_id, feed_id, lastFetchedAt)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: ListAllFeedPosts :many
SELECT p.*
FROM posts p
         INNER JOIN feeds f ON f.id = p.feedid
         INNER JOIN feedfollows ff ON ff.feed_id = f.id
WHERE ff.user_id = $1;
