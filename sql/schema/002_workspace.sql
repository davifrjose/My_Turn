-- +goose Up
CREATE TABLE workspaces (
  id   UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  display_name VARCHAR(255) NOT NULL,
  opening_time TIMESTAMP NOT NULL,
  closing_time TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE workspaces;