-- name: CreateUser :one
INSERT INTO users (id, created_at, email, name, password)
VALUES ($1, $2, $3, $4,$5)
RETURNING *;

-- name: SelectUserById :one
Select * from users where id = $1 LIMIT 1;
