package delivery

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateRoleDto struct {
		Name string
	}

	UpdateRoleDto struct {
		ID   uuid.UUID
		Name string
	}

	UpdateRolePermissionsDto struct {
		ID             uuid.UUID
		PermissionsIDs []uuid.UUID
	}

	RoleDto struct {
		ID               uuid.UUID
		Name             string
		CreationDate     time.Time
		LastUpdateDate   time.Time
		MarkAsDeleteDate *time.Time
		Permissions      []PermissionDto
	}

	RolesListDto struct {
		Count int
		Roles []RoleDto
	}
)
