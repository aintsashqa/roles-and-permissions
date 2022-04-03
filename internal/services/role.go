package services

import (
	"context"
	"time"

	"github.com/aintsashqa/roles-and-permissions/ent"
	"github.com/aintsashqa/roles-and-permissions/ent/permission"
	"github.com/aintsashqa/roles-and-permissions/ent/role"
	"github.com/aintsashqa/roles-and-permissions/internal/cerror"
	"github.com/aintsashqa/roles-and-permissions/internal/delivery"
	"github.com/google/uuid"
)

var (
	_ delivery.RoleService = &roleServiceImpl{}
)

type (
	roleServiceImpl struct {
		client *ent.Client
	}
)

func NewRoleServiceImpl(client *ent.Client) delivery.RoleService {
	return &roleServiceImpl{client: client}
}

func (srv *roleServiceImpl) Create(ctx context.Context, dto delivery.CreateRoleDto) (delivery.RoleDto, error) {
	roleResult, err := srv.client.Role.Create().SetName(dto.Name).Save(ctx)
	if err != nil {
		return delivery.RoleDto{}, cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	return delivery.RoleDto{
		ID:             roleResult.ID,
		Name:           roleResult.Name,
		CreationDate:   roleResult.CreationDate,
		LastUpdateDate: roleResult.LastUpdateDate,
	}, nil
}

func (srv *roleServiceImpl) Update(ctx context.Context, dto delivery.UpdateRoleDto) (delivery.RoleDto, error) {
	roleResult, err := srv.client.Role.Query().
		Where(role.MarkAsDeleteDateIsNil(), role.ID(dto.ID)).Only(ctx)
	if err != nil {
		return delivery.RoleDto{}, cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	roleResult, err = roleResult.Update().SetName(dto.Name).Save(ctx)
	if err != nil {
		return delivery.RoleDto{}, cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	return delivery.RoleDto{
		ID:             roleResult.ID,
		Name:           roleResult.Name,
		CreationDate:   roleResult.CreationDate,
		LastUpdateDate: roleResult.LastUpdateDate,
	}, nil
}

func (srv *roleServiceImpl) Delete(ctx context.Context, roleId uuid.UUID) error {
	_, err := srv.client.Role.Update().
		SetMarkAsDeleteDate(time.Now()).
		Where(role.MarkAsDeleteDateIsNil(), role.ID(roleId)).Save(ctx)
	if err != nil {
		return cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	return nil
}

func (srv *roleServiceImpl) GetRole(ctx context.Context, roleId uuid.UUID) (delivery.RoleDto, error) {
	roleResult, err := srv.client.Role.Query().
		WithPermissions(func(q *ent.PermissionQuery) {
			q.Select(permission.FieldID, permission.FieldName)
			q.Where(permission.MarkAsDeleteDateIsNil())
		}).
		Where(role.MarkAsDeleteDateIsNil(), role.ID(roleId)).Only(ctx)
	if err != nil {
		return delivery.RoleDto{}, cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	var permissions []delivery.PermissionDto
	for _, p := range roleResult.Edges.Permissions {
		temp := delivery.PermissionDto{
			ID:   p.ID,
			Name: p.Name,
		}

		permissions = append(permissions, temp)
	}

	return delivery.RoleDto{
		ID:             roleResult.ID,
		Name:           roleResult.Name,
		CreationDate:   roleResult.CreationDate,
		LastUpdateDate: roleResult.LastUpdateDate,
		Permissions:    permissions,
	}, nil
}

func (srv *roleServiceImpl) GetAvailableRolesList(ctx context.Context, includeDeleted bool) (delivery.RolesListDto, error) {
	query := srv.client.Role.Query()
	if !includeDeleted {
		query = query.Where(role.MarkAsDeleteDateIsNil())
	}

	rolesResult, err := query.All(ctx)
	if err != nil {
		return delivery.RolesListDto{}, cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	var roles []delivery.RoleDto
	for _, r := range rolesResult {
		temp := delivery.RoleDto{
			ID:               r.ID,
			Name:             r.Name,
			CreationDate:     r.CreationDate,
			LastUpdateDate:   r.LastUpdateDate,
			MarkAsDeleteDate: r.MarkAsDeleteDate,
		}

		roles = append(roles, temp)
	}

	return delivery.RolesListDto{
		Count: len(roles),
		Roles: roles,
	}, nil
}
