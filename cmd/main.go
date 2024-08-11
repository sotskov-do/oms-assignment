package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sotskov-do/oms-assignment/internal/config"
	"github.com/sotskov-do/oms-assignment/internal/controllers"
	"github.com/sotskov-do/oms-assignment/internal/controllers/bms"
	"github.com/sotskov-do/oms-assignment/internal/logger"
	"github.com/sotskov-do/oms-assignment/internal/service/apartments"
	"github.com/sotskov-do/oms-assignment/internal/service/buildings"
	"github.com/sotskov-do/oms-assignment/internal/storage/postgres"
)

var (
	app *fiber.App
	db  *postgres.PostgresDatabase
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	go start(ctx)
	<-ctx.Done()
	stop(ctx)
}

func start(ctx context.Context) {
	// ENV
	var err error
	if os.Getenv(config.IsLocal) == "" {
		err := godotenv.Load(config.ConfigPath)
		if err != nil {
			slog.Log(ctx, logger.LevelCritical, "error loading config", "error", err)
			os.Exit(1)
		}
	}

	// Logger
	l := logger.New(true)
	slog.SetDefault(l)

	// DB
	db, err = postgres.New(ctx, os.Getenv(config.PgURL))
	if err != nil {
		slog.Log(ctx, logger.LevelCritical, "can't create db", "error", err)
		os.Exit(1)
	}
	err = db.Ping(ctx)
	if err != nil {
		slog.Log(ctx, logger.LevelCritical, "can't ping db", "error", err)
		os.Exit(1)
	}

	// BMS
	apartmentsService := apartments.NewService(db)
	buildingsService := buildings.NewService(db)
	bms := bms.NewBuildingManagementSystem(apartmentsService, buildingsService)

	// App
	app = fiber.New()
	controllers.SetupRoutes(app, bms)
	go app.Listen(":3000")

	slog.Info("app started")
}

func stop(ctx context.Context) {
	slog.Info("shutting down")
	_ = app.ShutdownWithContext(ctx)
	_ = db.Stop(ctx)
}
