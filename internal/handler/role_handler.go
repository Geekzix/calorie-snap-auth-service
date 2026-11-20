package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Geekzix/calorie-snap-auth-service/internal/domain"
	"github.com/Geekzix/calorie-snap-auth-service/internal/service"
	"github.com/Geekzix/calorie-snap-commons/common/constant"
	"github.com/Geekzix/calorie-snap-commons/handler"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

type RoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func toRoleResponse(r *domain.Role) RoleResponse {
	return RoleResponse{
		ID:          r.ID.String(),
		Name:        r.Name,
		Description: r.Description,
	}
}

func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) error {
	var req RoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		return constant.InvalidInput("invalid request body")
	}

	role, err := h.svc.Create(r.Context(), req.Name, req.Description)
	if err != nil {
		return err
	}

	return handler.RespondJSON(w, http.StatusCreated, toRoleResponse(role))
}

func (h *RoleHandler) GetByID(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return constant.InvalidInput("invalid role id")
	}

	role, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		return err
	}

	return handler.RespondJSON(w, http.StatusOK, toRoleResponse(role))
}

func (h *RoleHandler) List(w http.ResponseWriter, r *http.Request) error {
	roles, err := h.svc.List(r.Context())
	if err != nil {
		return err
	}

	resp := make([]RoleResponse, 0, len(roles))
	for i := range roles {
		resp = append(resp, toRoleResponse(&roles[i]))
	}

	return handler.RespondJSON(w, http.StatusOK, resp)
}

func (h *RoleHandler) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return constant.InvalidInput("invalid role id")
	}

	var req RoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		return constant.InvalidInput("invalid request body")
	}

	role, err := h.svc.Update(r.Context(), id, req.Name, req.Description)
	if err != nil {
		return err
	}

	return handler.RespondJSON(w, http.StatusOK, toRoleResponse(role))
}

func (h *RoleHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return constant.InvalidInput("invalid role id")
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

