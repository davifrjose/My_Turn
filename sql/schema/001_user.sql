-- +goose Up
CREATE TABLE users (
  id   UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  email VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE users;