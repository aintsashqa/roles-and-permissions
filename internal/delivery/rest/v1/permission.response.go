package v1

import (
	"time"

	"github.com/google/uuid"
)

type (
	PermissionResponse struct {
		ID               uuid.UUID  `json:"id"`
		Name             string     `json:"name"`
		CreationDate     time.Time  `json:"creation_date"`
		LastUpdateDate   time.Time  `json:"last_update_date"`
		MarkAsDeleteDate *time.Time `json:"mark_as_delete_date,omitempty"`
	}

	PermissionsListResponse struct {
		Count       int                  `json:"count"`
		Permissions []PermissionResponse `json:"permissions"`
	}
)
