package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Lacky1234union/UrlShorter/internal/config"
	"github.com/Lacky1234union/UrlShorter/internal/lib/logger/sl"
	"github.com/Lacky1234union/UrlShorter/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad() // TODO: init config: cleanenv

	fmt.Println(cfg)

	log := setupLogger(cfg.Env)
	log.Info("Server Start", slog.String("env", cfg.Env))
	//
	//TODO: init logger: slog (logger)

	// TODO: init storage
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed init storage", sl.Err(err))
		os.Exit(1)
	}
	_ = storage
	//TODO: init roouter: chi, "chi render"
	//
	//TODO: run server
	id, err := storage.SaveURL("https://google.com", "google")
	if err != nil {
		log.Error("faled save url", sl.Err(err))
	}
	log.Info("saved url", slog.Int64("id", id))
	str, err := storage.GetURL("google")
	if err != nil {
		log.Error("faled get url", sl.Err(err))
	}
	log.Info("get url", slog.String("url", str))
	err = storage.DeleteURL("google")
	if err != nil {
		log.Error("faled delete url", sl.Err(err))
	}
	log.Info("delete url", slog.String("alias", "google"))
	str, err = storage.GetURL("google")
	if err != nil {
		log.Error("faled get url", sl.Err(err))
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envDev:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	case envProd:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
	}
	return log
}
