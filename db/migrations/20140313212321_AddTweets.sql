
-- +goose Up
CREATE TABLE tweets (
  Id serial not null primary key,
  CreatedAt text,
  TwitterId integer,
  Text text,
  School text
);


-- +goose Down
DROP TABLE tweets;
