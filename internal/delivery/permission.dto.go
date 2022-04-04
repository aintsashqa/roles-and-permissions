package delivery

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreatePermissionDto struct {
		Name string
	}

	PermissionDto struct {
		ID               uuid.UUID
		Name             string
		CreationDate     time.Time
		LastUpdateDate   time.Time
		MarkAsDeleteDate *time.Time
	}

	PermissionsListDto struct {
		Count       int
		Permissions []PermissionDto
	}
)
