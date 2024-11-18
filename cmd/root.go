package cmd

import (
	"log/slog"
	"os"

	"github.com/clipclock08/news-crud/cmd/server"
	"github.com/clipclock08/news-crud/internal/config"
)

func Exec() error {
	cfg, err := config.LoadConfigFromEnv()
	if err != nil {
		return err
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: cfg.Server.LogLevel,
	}))

	if err := server.Run(cfg, logger); err != nil {
		return err
	}

	return nil
}
