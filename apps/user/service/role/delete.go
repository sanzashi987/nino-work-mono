package roleService

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

func DeleteRole(ctx context.Context, roleId uint64) error {
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
