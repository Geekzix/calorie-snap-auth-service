package router

import (
	"log/slog"
	"net/http"

	"github.com/Geekzix/calorie-snap-auth-service/internal/handler"
	appmiddleware "github.com/Geekzix/calorie-snap-auth-service/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func New(logger *slog.Logger, db *gorm.DB) http.Handler {
	r := chi.NewRouter()

	// Chi Standard middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr)
	r.Use(middleware.Recoverer)

	// App middleware
	r.Use(appmiddleware.RequestLogger(logger))
	// http endpoints
	r.Get("/health", handler.Health)
	r.Get("/ready", handler.Ready)

	r.Route("/roles", roleRoutes(db, logger))

	return r
}
