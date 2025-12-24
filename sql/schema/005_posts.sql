-- +goose Up
CREATE TABLE Posts(
    Id UUID PRIMARY KEY,
    Title VARCHAR(64) NOT NULL,
    Description VARCHAR(64),
    PublishedAt TIMESTAMP NOT NULL,
    Link VARCHAR(255) NOT NULL,
    FeedId UUID REFERENCES Feeds(id) NOT NULL
);

-- +goose Down
DROP TABLE Posts;