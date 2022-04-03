package v1

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/aintsashqa/roles-and-permissions/internal/cerror"
	"github.com/google/uuid"
)

type (
	CreateRoleRequest struct {
		Name string `json:"name"`
	}

	UpdateRoleRequest struct {
		ID   uuid.UUID `json:"-"`
		Name string    `json:"name"`
	}
)

func (req *CreateRoleRequest) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	return nil
}

func (req *UpdateRoleRequest) Parse(r *http.Request) error {
	roleIdParam := chi.URLParam(r, "role_id")
	if len(roleIdParam) != 0 {
		roleId, err := uuid.Parse(roleIdParam)
		if err != nil {
			return cerror.Wrap(err, cerror.InternalComplexErrorType)
		}

		req.ID = roleId
	}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	return nil
}
