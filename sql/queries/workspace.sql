-- name: CreateWorkspaces :one
INSERT INTO workspaces (id, created_at, name, email, address, user_id, display_name, opening_time, closing_time)
VALUES ($1, $2, $3, $4,$5,$6,$7,$8,$9)
RETURNING *;
