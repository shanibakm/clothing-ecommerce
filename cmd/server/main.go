package main

import (
	"fmt"
	"os"

	"github.com/shanibakm/clothing-ecommerce/config"
	"github.com/shanibakm/clothing-ecommerce/internal/infrastructure/database"
	"github.com/shanibakm/clothing-ecommerce/internal/infrastructure/logger"
	"github.com/shanibakm/clothing-ecommerce/internal/router"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Configuration error:", err)
		os.Exit(1)
	}

	log, err := logger.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Logger error:", err)
		os.Exit(1)
	}
	defer log.Sync()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Database connection failed")// Avoid logging the DSN or database password.

	}

	r := router.New(db)

	log.Info("Application starting",
		zap.String("environment", cfg.AppEnv),
		zap.String("port", cfg.AppPort),
	)

	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("Application stopped unexpectedly")
	}
}
