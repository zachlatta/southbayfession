
-- +goose Up
ALTER TABLE tweets
  RENAME TO Tweet;


-- +goose Down
ALTER TABLE Tweet
  RENAME TO tweets;
