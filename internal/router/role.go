package router

import (
	"log/slog"

	"github.com/Geekzix/calorie-snap-auth-service/internal/handler"
	"github.com/Geekzix/calorie-snap-auth-service/internal/repository"
	"github.com/Geekzix/calorie-snap-auth-service/internal/service"
	global "github.com/Geekzix/calorie-snap-commons/handler"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func roleRoutes(db *gorm.DB, logger *slog.Logger) func(r chi.Router) {
	roleRepo := repository.NewRoleRepo(db)
	roleSvc := service.NewRoleSvc(roleRepo)
	roleHandler := handler.NewRoleHandler(roleSvc, logger)

	return func(r chi.Router) {
		r.Get("/", global.GlobalHandler(roleHandler.List))
		r.Post("/create", global.GlobalHandler(roleHandler.Create))
		r.Get("/{id}", global.GlobalHandler(roleHandler.GetByID))
		r.Put("/{id}", global.GlobalHandler(roleHandler.Update))
		r.Delete("/{id}", global.GlobalHandler(roleHandler.Delete))
	}
}
