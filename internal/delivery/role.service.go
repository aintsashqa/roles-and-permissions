package delivery

import (
	"context"

	"github.com/google/uuid"
)

type (
	RoleService interface {
		Create(ctx context.Context, dto CreateRoleDto) (RoleDto, error)
		Update(ctx context.Context, dto UpdateRoleDto) (RoleDto, error)
		Delete(ctx context.Context, roleId uuid.UUID) error
		GetRole(ctx context.Context, roleId uuid.UUID) (RoleDto, error)
		GetAvailableRolesList(ctx context.Context) (RolesListDto, error)
	}
)
