package server

import (
	"context"
	"errors"
	"github.com/clipclock08/news-crud/internal/handlers"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/clipclock08/news-crud/internal/app"
	"github.com/clipclock08/news-crud/internal/config"
	"github.com/clipclock08/news-crud/internal/db"
)

func Run(cfg config.Config, logger *slog.Logger) error {
	/*pg, err := mongo.New(cfg)
	if err != nil {
		return err
	}*/
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	dbClient := db.Connect("mongodb://localhost:27017")
	app, err := app.New(app.WithMongoRepo(dbClient))
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.NewHomeHandler().ServeHTTP)
	server := &http.Server{
		Addr:         cfg.Server.Address,
		Handler:      mux,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	go func() {
		switch cfg.Server.Environment {
		case config.Production:
			log.Printf("starting server on %s environment", cfg.Server.Environment)
		default:
			logger.Info("server",
				slog.String("environment", string(cfg.Server.Environment)),
				slog.String("listening on", cfg.Server.Address),
			)
		}
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(err.Error())
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.WriteTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
