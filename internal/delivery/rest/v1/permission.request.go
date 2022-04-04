package v1

import (
	"encoding/json"
	"net/http"

	"github.com/aintsashqa/roles-and-permissions/internal/cerror"
)

type (
	CreatePermissionRequest struct {
		Name string `json:"name"`
	}
)

func (req *CreatePermissionRequest) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	return nil
}
