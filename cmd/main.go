package main

import (
	"log/slog"
	"url-shortner/internal/config"
	"url-shortner/internal/lib/logger"
	"url-shortner/internal/lib/logger/sl"
	"url-shortner/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	log = log.With(slog.String("env", cfg.Env))
	log.Info("initializing server", slog.String("address", cfg.Address))

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to initialize storage", sl.Err(err))
	}

	_ = storage
}
