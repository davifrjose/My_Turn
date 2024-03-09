package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/davifrjose/My_Turn/internal/adapter/config"
	"github.com/davifrjose/My_Turn/internal/adapter/handler/http"
	"github.com/davifrjose/My_Turn/internal/adapter/logger"
	"github.com/davifrjose/My_Turn/internal/adapter/storage/postgres"
	"github.com/davifrjose/My_Turn/internal/adapter/storage/postgres/repository"
	"github.com/davifrjose/My_Turn/internal/core/service"
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

	serviceTypeRepository := repository.NewServiceTypeRepository(db)
	serviceTypeService := service.NewServiceType(serviceTypeRepository)
	serviceTypeHandler := http.NewServiceTypeHandler(serviceTypeService)

	router, err := http.NewRouter(
		config.HTTP,
		*serviceTypeHandler,
	)
	if err != nil {
		slog.Info("Error initializing router", "error", err)
		os.Exit(1)
	}

	listenAddress := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	slog.Info("Starting http server", "listen address", config.HTTP.Port)
	err = router.Serve(listenAddress)
	if err != nil {
		slog.Error("Error starting http server", "error", err)
		os.Exit(1)
	}
}
