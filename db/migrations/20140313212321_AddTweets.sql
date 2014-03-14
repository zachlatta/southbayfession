
-- +goose Up
CREATE TABLE tweets (
  Id integer not null primary key autoincrement,
  CreatedAt text,
  TwitterId integer,
  Text text,
  School text
);


-- +goose Down
DROP TABLE tweets;
