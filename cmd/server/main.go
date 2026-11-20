package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Geekzix/calorie-snap-auth-service/internal/config"
	"github.com/Geekzix/calorie-snap-auth-service/internal/router"
	"github.com/Geekzix/calorie-snap-commons/utils"
	"github.com/go-chi/chi/v5"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.Error("failed to load configuration", "error", err)
		os.Exit(1)
	}

	handler := router.New(logger)

	basePath := utils.GetEnvString("BASE_URL_PATH", "/auth")
	apiHandler := chi.NewRouter()

	apiHandler.Mount(basePath, handler)

	server := &http.Server{
		Addr:         cfg.HttpAddr,
		Handler:      apiHandler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	serverErrors := make(chan error, 1)

	go func() {
		logger.Info("Auth server is running", "address", cfg.HttpAddr)
		serverErrors <- server.ListenAndServe()
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("http server failed", "error", err)
			os.Exit(1)
		}
	case <-ctx.Done():
		logger.Info("shutdown signal received")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("graceful shutdown failed", "error", err)

		if closeErr := server.Close(); closeErr != nil {
			logger.Error("server close failed", "error", closeErr)
		}

		os.Exit(1)
	}

	logger.Info("auth service stopped")

}
