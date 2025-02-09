package service

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	userService "github.com/sanzashi987/nino-work/apps/user/service/user"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

type RoleServiceWeb struct{}

var RoleServiceWebImpl *RoleServiceWeb = &RoleServiceWeb{}

// CreateRoleRequest 创建角色请求参数
type CreateRoleRequest struct {
	RoleName        string   `json:"role_name" binding:"required"`
	RoleCode        string   `json:"role_code" binding:"required"`
	RoleDescription string   `json:"role_description"`
	PermissionIds   []uint64 `json:"permission_ids"`
}

// 创建角色
func (r *RoleServiceWeb) CreateRole(ctx context.Context, userId uint64, payload CreateRoleRequest) error {
	user, tx, err := userService.GetUserRolePermission(ctx, userId)
	if err != nil {
		return err
	}

	if len(user.Roles) == 0 {
		return errors.New("user does not have any admin permission")
	}
	if len(payload.PermissionIds) == 0 {
		return errors.New("no permission to bind")
	}

	tx = tx.Begin()

	// 创建角色
	newRole := &model.RoleModel{
		Name:        payload.RoleName,
		Code:        payload.RoleCode,
		Description: payload.RoleDescription,
	}

	if err := tx.Create(newRole).Error; err != nil {
		tx.Rollback()
		return err
	}

	permissions := []*model.PermissionModel{}

	for _, pid := range payload.PermissionIds {
		permission := &model.PermissionModel{}
		permission.Id = pid
		permissions = append(permissions, permission)
	}

	if err := tx.Model(newRole).Association("Permissions").Replace(permissions); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// 获取角色详情
func (r *RoleServiceWeb) GetRoleDetail(ctx context.Context, roleId uint64) (*model.RoleModel, error) {

	role := model.RoleModel{}
	role.Id = roleId

	if err := dao.FindRolesWithPermissions(db.NewTx(ctx), &role); err != nil {
		return nil, err
	}

	return &role, nil
}

type UpdateRoleRequest struct {
	RoleName        string   `json:"role_name"`
	RoleId          uint64   `json:"role_id" binding:"required"`
	RoleDescription string   `json:"role_description"`
	PermissionIds   []uint64 `json:"permission_ids"`
}

// 更新角色
func (r *RoleServiceWeb) UpdateRole(ctx context.Context, payload UpdateRoleRequest) error {
	tx := db.NewTx(ctx).Begin()
	role := &model.RoleModel{}

	if payload.RoleName != "" {
		role.Name = payload.RoleName
	}
	if payload.RoleDescription != "" {
		role.Description = payload.RoleDescription
	}

	role.Id = payload.RoleId

	if err := tx.Updates(role).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新权限关联
	if len(payload.PermissionIds) > 0 {
		permissions := make([]model.PermissionModel, 0)
		for _, pid := range payload.PermissionIds {
			permission := model.PermissionModel{}
			permission.Id = pid
			permissions = append(permissions, permission)
		}

		if err := tx.Model(role).Association("Permissions").Replace(permissions); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

// 删除角色
func (r *RoleServiceWeb) DeleteRole(ctx context.Context, roleId uint64) error {
	tx := db.NewTx(ctx).Begin()

	tx.Begin()

	role := &model.RoleModel{}
	role.Id = roleId

	// 先清除角色关联的权限
	if err := tx.Model(role).Association("Permissions").Clear(); err != nil {
		tx.Rollback()
		return err
	}

	// 再删除角色本身
	if err := tx.Delete(role).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// 模糊搜索角色
func (r *RoleServiceWeb) SuggestRoles(ctx context.Context, keyword string) ([]model.RoleModel, error) {
	tx := db.NewTx(ctx)

	var roles []model.RoleModel

	if err := tx.
		Where("name LIKE ? OR code LIKE ?", "%"+keyword+"%", "%"+keyword+"%").
		Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

type RoleMeta struct {
	Name        string             `json:"name"`
	Code        string             `json:"code"`
	Permissions []*shared.EnumMeta `json:"permissions"`
}

type ListRolesResponse struct {
	Roles []*RoleMeta `json:"roles"`
	shared.PaginationResponse
}

// 列出角色
func (r *RoleServiceWeb) ListRoles(ctx context.Context, userId uint64, payload *shared.PaginationRequest) (*ListRolesResponse, error) {
	result, err := userService.GetUserAdmins(ctx, userId)
	if err != nil {
		return nil, err
	}

	tx := result.Tx

	if !result.HasAnyAdmin() {
		return &ListRolesResponse{
			Roles: []*RoleMeta{},
		}, nil
	}

	appIds := result.GetAllAppIds()

	var roles []model.RoleModel
	var totalCount int64

	offset := (payload.Page - 1) * payload.Size

	var permissions []*model.PermissionModel

	if err := tx.Where("permissions.app_id IN ?", appIds).Find(&permissions).Error; err != nil {
		return nil, err
	}

	permissionIds := make([]uint64, len(permissions))
	for i, permission := range permissions {
		permissionIds[i] = permission.Id
	}

	var roleIds = []*struct {
		RoleId uint64 `json:"role_model_id"`
	}
	subQuery := tx.Table("role_permissions").Select("role_model_id", "permission_model_id").Where("permission_model_id  IN ? ", permissionIds).Scan(&roles)

	// if err :=tx.Model(&model.RoleModel{}).

	roleMetas := make([]*RoleMeta, len(roles))
	for i, role := range roles {
		roleMetas[i] = &RoleMeta{
			Name: role.Name,
			Code: role.Code,
		}
	}

	return &ListRolesResponse{
		Roles: roleMetas,
		PaginationResponse: shared.PaginationResponse{
			PageIndex:   payload.Page,
			PageSize:    payload.Size,
			RecordTotal: int(totalCount),
		},
	}, nil
}
