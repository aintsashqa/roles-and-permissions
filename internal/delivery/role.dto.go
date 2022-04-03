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
