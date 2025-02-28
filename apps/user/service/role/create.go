package roleService

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	userService "github.com/sanzashi987/nino-work/apps/user/service/user"
)

// CreateRoleRequest 创建角色请求参数
type CreateRoleRequest struct {
	Name          string   `json:"name" binding:"required"`
	Code          string   `json:"code" binding:"required"`
	Description   string   `json:"description"`
	PermissionIds []uint64 `json:"permission_ids"`
}

func CreateRole(ctx context.Context, userId uint64, payload CreateRoleRequest) error {
	user, tx, err := userService.GetUserRolePermission(ctx, userId)
	if err != nil {
		return err
	}

	if len(user.Roles) == 0 {
		return errNopermission
	}
	if len(payload.PermissionIds) == 0 {
		return errEmptyPermission
	}

	tx = tx.Begin()

	// 创建角色
	newRole := &model.RoleModel{
		Name:        payload.Name,
		Code:        payload.Code,
		Description: payload.Description,
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
