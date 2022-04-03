package services

import (
	"context"

	"github.com/aintsashqa/roles-and-permissions/ent"
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
		Where(role.ID(dto.ID)).Only(ctx)
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
	_, err := srv.client.Role.Delete().
		Where(role.ID(roleId)).Exec(ctx)
	if err != nil {
		return cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	return nil
}

func (srv *roleServiceImpl) GetRole(ctx context.Context, roleId uuid.UUID) (delivery.RoleDto, error) {
	roleResult, err := srv.client.Role.Query().
		Where(role.ID(roleId)).Only(ctx)
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

func (srv *roleServiceImpl) GetAvailableRolesList(ctx context.Context) (delivery.RolesListDto, error) {
	rolesResult, err := srv.client.Role.Query().
		Where().All(ctx)
	if err != nil {
		return delivery.RolesListDto{}, cerror.Wrap(err, cerror.InternalComplexErrorType)
	}

	var roles []delivery.RoleDto
	for _, r := range rolesResult {
		temp := delivery.RoleDto{
			ID:             r.ID,
			Name:           r.Name,
			CreationDate:   r.CreationDate,
			LastUpdateDate: r.LastUpdateDate,
		}

		roles = append(roles, temp)
	}

	return delivery.RolesListDto{
		Count: len(roles),
		Roles: roles,
	}, nil
}
