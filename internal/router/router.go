package router

import (
	"log/slog"
	"net/http"

	"github.com/Geekzix/calorie-snap-auth-service/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(logger *slog.Logger) http.Handler {
	r := chi.NewRouter()

	// Chi Standard middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr)
	r.Use(middleware.Recoverer)

	// App middleware

	// http endpoints
	r.Get("/health", handler.Health)
	r.Get("/ready", handler.Ready)

	return r
}
