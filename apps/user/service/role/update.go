package roleService

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type UpdateRoleRequest struct {
	Id            uint64   `json:"id" binding:"required"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	PermissionIds []uint64 `json:"permission_ids"`
}

// 更新角色
func UpdateRole(ctx context.Context, payload UpdateRoleRequest) error {
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
