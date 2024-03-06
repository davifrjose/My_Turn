package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/davifrjose/My_Turn/internal/adapter/config"
	"github.com/davifrjose/My_Turn/internal/adapter/logger"
	"github.com/davifrjose/My_Turn/internal/adapter/storage/postgres"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.New()
	if err != nil {
		slog.Error("Error loafing env vars", "error", err)
		os.Exit(1)
	}

	logger.Set(config.App)
	slog.Info("Starting the application", "app", config.App.Name, "Env", config.App.Env)

	ctx := context.Background()
	db, err := postgres.New(ctx, config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Successfully connected to the database", "db", config.DB.ConnectionUrl)
}
