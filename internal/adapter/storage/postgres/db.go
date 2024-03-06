package postgres

import (
	"context"

	"github.com/davifrjose/My_Turn/internal/adapter/config"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Db struct {
	*pgxpool.Pool
	url string
}

func New(ctx context.Context, config *config.DB) (*Db, error) {
	db, err := pgxpool.New(ctx, config.ConnectionUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &Db{
		db,
		config.ConnectionUrl,
	}, nil
}

func (db *Db) Close() {
	db.Pool.Close()
}

func (db *Db) ErrorCode(err error) string {
	pgError := err.(*pgconn.PgError)

	return pgError.Code
}
