-- +goose Up
CREATE TABLE Users(
    id UUID PRIMARY KEY ,
    createdAt timestamp NOT NULL ,
    updatedAt timestamp NOT NULL ,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE Users;