package roleService

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

var errNopermission = errors.New("user does not have any admin permission")
var errEmptyPermission = errors.New("no permission to bind")

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
	Id            uint64   `json:"id" binding:"required"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	PermissionIds []uint64 `json:"permission_ids"`
}

// 更新角色
func (r *RoleServiceWeb) UpdateRole(ctx context.Context, payload UpdateRoleRequest) error {
	tx := db.NewTx(ctx).Begin()
	role := &model.RoleModel{}

	if payload.Name != "" {
		role.Name = payload.Name
	}
	if payload.Description != "" {
		role.Description = payload.Description
	}

	role.Id = payload.Id

	if err := tx.Updates(role).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新权限关联
	if len(payload.PermissionIds) > 0 {
		permissions := make([]*model.PermissionModel, 0)
		for _, pid := range payload.PermissionIds {
			permission := &model.PermissionModel{}
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

type RoleMeta struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	// Permissions []*shared.EnumMeta `json:"permissions"`
}

type ListRolesResponse struct {
	Data []*RoleMeta `json:"data"`
	shared.PaginationResponse
}

func (r *RoleServiceWeb) ListRoles(ctx context.Context, userId uint64, payload *shared.PaginationRequest) (*ListRolesResponse, error) {
	result, err := userService.GetUserAdmins(ctx, userId)
	if err != nil {
		return nil, err
	}

	tx := result.Tx

	if !result.HasAnyAdmin() {
		return &ListRolesResponse{
			Data: []*RoleMeta{},
		}, nil
	}

	appIds := result.GetAllAppIds()

	var roles []model.RoleModel
	var totalCount int64

	filteredPermissions := tx.Table("permissions").Where("app_id IN ?", appIds)

	subQuery := tx.Table("(?) as filtered_permissions", filteredPermissions).
		Select("DISTINCT role_permissions.role_model_id").
		Joins("LEFT JOIN role_permissions ON role_permissions.permission_model_id = filtered_permissions.id")

	query := tx.Table("(?) AS b", subQuery).
		Select("b.role_model_id, roles.*").
		Joins("LEFT JOIN roles ON b.role_model_id = roles.id").Where("roles.delete_time IS NULL")

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	if err := query.Order("roles.id DESC").Scopes(db.Paginate(payload.Page, payload.Size)).Scan(&roles).Error; err != nil {
		return nil, err
	}

	roleMetas := make([]*RoleMeta, len(roles))
	for i, role := range roles {
		roleMetas[i] = &RoleMeta{
			Id:   role.Id,
			Name: role.Name,
			Code: role.Code,
		}
	}

	return &ListRolesResponse{
		Data: roleMetas,
		PaginationResponse: shared.PaginationResponse{
			PageIndex:   payload.Page,
			PageSize:    payload.Size,
			RecordTotal: int(totalCount),
		},
	}, nil
}
