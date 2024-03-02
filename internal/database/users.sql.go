// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, email, name, password)
VALUES ($1, $2, $3, $4,$5)
RETURNING id, created_at, email, name, password
`

type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Email     string
	Name      string
	Password  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.Email,
		arg.Name,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Email,
		&i.Name,
		&i.Password,
	)
	return i, err
}
