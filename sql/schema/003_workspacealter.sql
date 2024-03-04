-- +goose Up
ALTER TABLE IF EXISTS workspaces
ADD COLUMN IF NOT EXISTS logo VARCHAR(255),
ADD COLUMN IF NOT EXISTS description text;
-- +goose Down
ALTER TABLE workspaces
DROP COLUMN logo, description;

