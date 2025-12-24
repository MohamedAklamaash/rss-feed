-- +goose Up

ALTER TABLE Users ADD COLUMN api_key varchar(64) UNIQUE NOT NULL DEFAULT (encode(sha256(random()::text::bytea),'hex'));

-- +goose Down
ALTER TABLE Users DROP COLUMN api_key;