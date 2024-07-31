package main

import (
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/valitovgaziz/url-shortener/internal/config"
	mwLogger "github.com/valitovgaziz/url-shortener/internal/http-server/middleware/logger"
	"github.com/valitovgaziz/url-shortener/internal/lib/logger/handlers/slogpretty"
	"github.com/valitovgaziz/url-shortener/internal/lib/logger/sl"
	"github.com/valitovgaziz/url-shortener/internal/storage/sqlite"
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

	// init storage
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage %w", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// TODO: run server

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	// setup logger
	switch env {
	case envLocal:
		log = setupPrettySlog(slog.LevelDebug)
	case envDev:
		log = setupPrettySlog(slog.LevelDebug)
	case envProd:
		log = setupPrettySlog(slog.LevelInfo)
	}
	return log
}

func setupPrettySlog(level slog.Level) *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
