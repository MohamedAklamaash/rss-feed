-- +goose Up
CREATE TABLE FeedFollows(
    id UUID PRIMARY KEY ,
    createdAt timestamp NOT NULL ,
    updatedAt timestamp NOT NULL ,
    user_id UUID NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES Feeds(id) ON DELETE CASCADE,
    UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE FeedFollows;