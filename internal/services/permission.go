package services

import (
	"context"
	"time"

	"github.com/aintsashqa/roles-and-permissions/ent"
	"github.com/aintsashqa/roles-and-permissions/ent/permission"
	"github.com/aintsashqa/roles-and-permissions/internal/cerror"
	"github.com/aintsashqa/roles-and-permissions/internal/delivery"
	"github.com/google/uuid"
)

type (
	permissionServiceImpl struct {
		client *ent.Client
	}
)

func NewPermissionServiceImpl(client *ent.Client) delivery.PermissionService {
	return &permissionServiceImpl{client: client}
}

func (srv *permissionServiceImpl) Create(ctx context.Context, dto delivery.CreatePermissionDto) (delivery.PermissionDto, error) {
	permissionResult, err := srv.client.Permission.Create().SetName(dto.Name).Save(ctx)
	if err != nil {
		return delivery.PermissionDto{}, cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	return delivery.PermissionDto{
		ID:             permissionResult.ID,
		Name:           permissionResult.Name,
		CreationDate:   permissionResult.CreationDate,
		LastUpdateDate: permissionResult.LastUpdateDate,
	}, nil
}

func (srv *permissionServiceImpl) Delete(ctx context.Context, permissionId uuid.UUID) error {
	_, err := srv.client.Permission.Update().
		SetMarkAsDeleteDate(time.Now()).
		Where(permission.MarkAsDeleteDateIsNil(), permission.ID(permissionId)).Save(ctx)
	if err != nil {
		return cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	return nil
}

func (srv *permissionServiceImpl) GetAvailablePermissionsList(ctx context.Context, includeDeleted bool) (delivery.PermissionsListDto, error) {
	query := srv.client.Permission.Query()
	if !includeDeleted {
		query = query.Where(permission.MarkAsDeleteDateIsNil())
	}

	permissionsResult, err := query.All(ctx)
	if err != nil {
		return delivery.PermissionsListDto{}, cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	var permissions []delivery.PermissionDto
	for _, r := range permissionsResult {
		temp := delivery.PermissionDto{
			ID:               r.ID,
			Name:             r.Name,
			CreationDate:     r.CreationDate,
			LastUpdateDate:   r.LastUpdateDate,
			MarkAsDeleteDate: r.MarkAsDeleteDate,
		}

		permissions = append(permissions, temp)
	}

	return delivery.PermissionsListDto{
		Count:       len(permissions),
		Permissions: permissions,
	}, nil
}
