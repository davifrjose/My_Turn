package postgres

import (
	"context"
	"embed"

	"github.com/davifrjose/My_Turn/internal/adapter/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var migrationsFs embed.FS

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

func (db *Db) Migrate() error {
	driver, err := iofs.New(migrationsFs, "migrations")
	if err != nil {
		return err
	}

	migrations, err := migrate.NewWithSourceInstance("iofs", driver, db.url)
	if err != nil {
		return err
	}

	err = migrations.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
