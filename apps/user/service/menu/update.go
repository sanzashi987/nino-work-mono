package menuService

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	userService "github.com/sanzashi987/nino-work/apps/user/service/user"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type UpdateMenuRequest struct {
	Id          uint64    `json:"id" binding:"required"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Order       *int      `json:"order"`
	Status      *int      `json:"status"`
	Path        string    `json:"path"`
	Roles       *[]uint64 `json:"roles"`
}

func Update(ctx context.Context, userId uint64, req *UpdateMenuRequest) error {
	tx := db.NewTx(ctx)
	exsitMenu := model.MenuModel{}
	if err := tx.Preload("Roles", "menu_model_id = ?", req.Id).Where("id = ?", req.Id).Find(&exsitMenu).Error; err != nil {
		return err
	}

	toUpdate := model.MenuModel{}
	if req.Name != "" {
		toUpdate.Name = req.Name
	}
	if req.Description != "" {
		toUpdate.Description = req.Description
	}
	if req.Order != nil {
		toUpdate.Order = *req.Order
	}
	if req.Status != nil {
		toUpdate.Status = *req.Status
	}
	if req.Path != "" {
		toUpdate.Path = req.Path
	}

	if req.Roles != nil {
		nextRoleIds := *req.Roles

		roles := make([]db.IGetId, len(exsitMenu.Roles))
		for i, role := range exsitMenu.Roles {
			roles[i] = role
		}
		exsitRoles := db.ToIdList(roles)

		nextSet, err := dao.FindAllPermissionsWithRoleIds(tx, nextRoleIds)
		if err != nil {
			return err
		}

		exsitSet, err := dao.FindAllPermissionsWithRoleIds(tx, exsitRoles)
		if err != nil {
			return err
		}

		_, changed := exsitSet.Diff(nextSet)
		result, err := userService.GetUserAdmins(ctx, userId)
		if err != nil {
			return err
		}

		hasPermission, err := result.ToPermissionSet()
		if err != nil {
			return err
		}

		if !hasPermission.IsStrictlyContains(changed) {
			return userService.ErrOutsidepermission
		}
	}

	tx = tx.Begin()
	if err := tx.Model(&exsitMenu).Updates(&toUpdate).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil

}
