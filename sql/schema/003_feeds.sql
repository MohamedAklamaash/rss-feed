-- +goose Up
CREATE TABLE Feeds(
    id UUID PRIMARY KEY ,
    createdAt timestamp NOT NULL ,
    updatedAt timestamp NOT NULL ,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) on DELETE CASCADE
);

-- +goose Down
DROP TABLE Feeds;