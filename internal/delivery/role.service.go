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
		GetAvailableRolesList(ctx context.Context, includeDeleted bool) (RolesListDto, error)

		AttachPermissions(ctx context.Context, dto UpdateRolePermissionsDto) error
		DetachPermissions(ctx context.Context, dto UpdateRolePermissionsDto) error
	}
)
