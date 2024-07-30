package main

import (
	"log/slog"
	"os"

	"github.com/valitovgaziz/url-shortener/internall/config"
	"github.com/valitovgaziz/url-shortener/internall/lib/logger/sl"
	"github.com/valitovgaziz/url-shortener/internall/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func init() {
	os.Setenv("CONFIG_PATH", "D:\\GoLang\\url-shortener\\config\\local.yaml")
}

func main() {

	// load config
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")


	storage, err := sqlite.New(cfg.StoragePath)
	if err!= nil {
		log.Error("failed to init storage %w", sl.Err(err))
		os.Exit(1)
	}
	_ = storage
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	// setup logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
