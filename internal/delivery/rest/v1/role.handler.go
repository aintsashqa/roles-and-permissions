package v1

import (
	"encoding/json"
	"net/http"

	"github.com/aintsashqa/roles-and-permissions/internal/cerror"
	"github.com/aintsashqa/roles-and-permissions/internal/delivery"
	pkg "github.com/aintsashqa/roles-and-permissions/pkg/http"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type (
	roleHandler struct {
		service delivery.RoleService
	}
)

func RegisterRoleRoutes(r chi.Router, container *delivery.Container) {
	handler := roleHandler{service: container.RoleService}

	r.Route("/role", func(r chi.Router) {
		r.Post("/", pkg.ErrorHandler(handler.Create).ServeHTTP)
		r.Get("/", pkg.ErrorHandler(handler.GetRolesList).ServeHTTP)

		r.Put("/{role_id}", pkg.ErrorHandler(handler.Update).ServeHTTP)
		r.Get("/{role_id}", pkg.ErrorHandler(handler.GetRole).ServeHTTP)
		r.Delete("/{role_id}", pkg.ErrorHandler(handler.Delete).ServeHTTP)
	})
}

func (handler *roleHandler) Create(w http.ResponseWriter, r *http.Request) error {
	var req CreateRoleRequest
	if err := req.Parse(r); err != nil {
		return err
	}

	dto := delivery.CreateRoleDto{Name: req.Name}
	result, err := handler.service.Create(r.Context(), dto)
	if err != nil {
		return err
	}

	resp := RoleResponse{
		ID:             result.ID,
		Name:           result.Name,
		CreationDate:   result.CreationDate,
		LastUpdateDate: result.LastUpdateDate,
	}

	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(resp)
}

func (handler *roleHandler) Update(w http.ResponseWriter, r *http.Request) error {
	var req UpdateRoleRequest
	if err := req.Parse(r); err != nil {
		return err
	}

	dto := delivery.UpdateRoleDto{
		ID:   req.ID,
		Name: req.Name,
	}

	result, err := handler.service.Update(r.Context(), dto)
	if err != nil {
		return err
	}

	resp := RoleResponse{
		ID:             result.ID,
		Name:           result.Name,
		CreationDate:   result.CreationDate,
		LastUpdateDate: result.LastUpdateDate,
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func (handler *roleHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	roleIdParam := chi.URLParam(r, "role_id")
	if len(roleIdParam) == 0 {
		return cerror.New("required role id cannot be empty string", cerror.InternalComplexErrorType)
	}

	roleId, err := uuid.Parse(roleIdParam)
	if err != nil {
		return cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	if err := handler.service.Delete(r.Context(), roleId); err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (handler *roleHandler) GetRole(w http.ResponseWriter, r *http.Request) error {
	roleIdParam := chi.URLParam(r, "role_id")
	if len(roleIdParam) == 0 {
		return cerror.New("required role id cannot be empty string", cerror.InternalComplexErrorType)
	}

	roleId, err := uuid.Parse(roleIdParam)
	if err != nil {
		return cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	result, err := handler.service.GetRole(r.Context(), roleId)
	if err != nil {
		return err
	}

	resp := RoleResponse{
		ID:             result.ID,
		Name:           result.Name,
		CreationDate:   result.CreationDate,
		LastUpdateDate: result.LastUpdateDate,
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func (handler *roleHandler) GetRolesList(w http.ResponseWriter, r *http.Request) error {
	result, err := handler.service.GetAvailableRolesList(r.Context())
	if err != nil {
		return err
	}

	var roles []RoleResponse
	for _, r := range result.Roles {
		temp := RoleResponse{
			ID:             r.ID,
			Name:           r.Name,
			CreationDate:   r.CreationDate,
			LastUpdateDate: r.LastUpdateDate,
		}

		roles = append(roles, temp)
	}

	resp := RolesListResponse{
		Count: result.Count,
		Roles: roles,
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}
