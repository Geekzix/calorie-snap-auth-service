package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Geekzix/calorie-snap-auth-service/internal/service"
	"github.com/Geekzix/calorie-snap-commons/common/constant"
	"github.com/Geekzix/calorie-snap-commons/handler"
)

type RoleHandler struct {
	svc    *service.RoleSvc
	logger *slog.Logger
}

func NewRoleHandler(svc *service.RoleSvc, logger *slog.Logger) *RoleHandler {
	return &RoleHandler{
		svc:    svc,
		logger: logger,
	}
}

type CreateRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleResponse struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) error {
	var req CreateRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", r.Body)
		return constant.InvalidInput("invalid request body")
	}

	role, err := h.svc.Create(r.Context(), req.Name, req.Description)
	if err != nil {
		return err
	}

	response := RoleResponse{
		ID:          role.ID.String(),
		Name:        role.Name,
		Description: role.Description,
	}

	return handler.RespondJSON(w, http.StatusCreated, response)

}
