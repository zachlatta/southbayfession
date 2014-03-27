
-- +goose Up
ALTER TABLE Tweet
  ALTER COLUMN TwitterId SET DATA TYPE bigint;


-- +goose Down
ALTER TABLE Tweet
  ALTER COLUMN TwitterId SET DATA TYPE integer;
