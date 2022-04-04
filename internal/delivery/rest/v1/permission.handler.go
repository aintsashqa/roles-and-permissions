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
	permissionHandler struct {
		service delivery.PermissionService
	}
)

func RegisterPermissionRoutes(r chi.Router, container *delivery.Container) {
	handler := permissionHandler{service: container.PermissionService}

	r.Route("/permission", func(r chi.Router) {
		r.Post("/", pkg.ErrorHandler(handler.Create).ServeHTTP)
		r.Get("/", pkg.ErrorHandler(handler.GetPermissionsList).ServeHTTP)

		r.Delete("/{permission_id}", pkg.ErrorHandler(handler.Delete).ServeHTTP)
	})
}

func (handler *permissionHandler) Create(w http.ResponseWriter, r *http.Request) error {
	var req CreatePermissionRequest
	if err := req.Parse(r); err != nil {
		return err
	}

	dto := delivery.CreatePermissionDto{Name: req.Name}
	result, err := handler.service.Create(r.Context(), dto)
	if err != nil {
		return err
	}

	resp := PermissionResponse{
		ID:             result.ID,
		Name:           result.Name,
		CreationDate:   result.CreationDate,
		LastUpdateDate: result.LastUpdateDate,
	}

	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(resp)
}

func (handler *permissionHandler) GetPermissionsList(w http.ResponseWriter, r *http.Request) error {
	includeDeleted := r.URL.Query().Has("include_deleted")
	result, err := handler.service.GetAvailablePermissionsList(r.Context(), includeDeleted)
	if err != nil {
		return err
	}

	var permissions []PermissionResponse
	for _, r := range result.Permissions {
		temp := PermissionResponse{
			ID:               r.ID,
			Name:             r.Name,
			CreationDate:     r.CreationDate,
			LastUpdateDate:   r.LastUpdateDate,
			MarkAsDeleteDate: r.MarkAsDeleteDate,
		}

		permissions = append(permissions, temp)
	}

	resp := PermissionsListResponse{
		Count:       result.Count,
		Permissions: permissions,
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func (handler *permissionHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	permissionIdParam := chi.URLParam(r, "permission_id")
	if len(permissionIdParam) == 0 {
		return cerror.New("required permission id cannot be empty string", cerror.InternalComplexErrorType)
	}

	permissionId, err := uuid.Parse(permissionIdParam)
	if err != nil {
		return cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	if err := handler.service.Delete(r.Context(), permissionId); err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
