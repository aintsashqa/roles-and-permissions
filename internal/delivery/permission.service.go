package delivery

import (
	"context"

	"github.com/google/uuid"
)

type (
	PermissionService interface {
		Create(ctx context.Context, dto CreatePermissionDto) (PermissionDto, error)
		Delete(ctx context.Context, permissionId uuid.UUID) error
		GetAvailablePermissionsList(ctx context.Context, includeDeleted bool) (PermissionsListDto, error)
	}
)
