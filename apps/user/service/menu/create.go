package menuService

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	userService "github.com/sanzashi987/nino-work/apps/user/service/user"
)

type CreateMenuRequest struct {
	Name   string   `json:"name" binding:"required"`
	Code   string   `json:"code" binding:"required"`
	Type   uint8    `json:"type" binding:"required"`
	Order  int      `json:"order"`
	Status int      `json:"status"`
	Path   string   `json:"path" binding:"required"`
	Roles  []uint64 `json:"roles"`
}

func Create(ctx context.Context, userId uint64, req *CreateMenuRequest) error {
	result, err := userService.GetUserAdmins(ctx, userId)
	if err != nil {
		return err
	}

	if !result.HasAnyAdmin() {
		return userService.ErrNopermission
	}

	hasPermission, err := result.ToPermissionSet()
	if err != nil {
		return err
	}

	bindRole := len(req.Roles) > 0

	if bindRole {
		tryToUsePermission, err := dao.FindAllPermissionsWithRoleIds(result.Tx, req.Roles)
		if err != nil {
			return err
		}
		if !hasPermission.IsStrictlyContains(tryToUsePermission) {
			return userService.ErrOutsidepermission
		}
	}
	toCreateMenu := model.MenuModel{
		Name:   req.Name,
		Code:   req.Code,
		Type:   model.MenuType(req.Type),
		Order:  req.Order,
		Status: req.Status,
		Path:   req.Path,
	}

	tx := result.Tx.Begin()
	if err := tx.Create(&toCreateMenu).Error; err != nil {
		tx.Rollback()
		return err
	}

	if bindRole {
		toBind := make([]*model.RoleModel, len(req.Roles))
		for index, id := range req.Roles {
			role := model.RoleModel{}
			role.Id = id
			toBind[index] = &role
		}

		if err := tx.Model(&toCreateMenu).Association("Roles").Replace(&toBind); err != nil {
			tx.Rollback()
			return err
		}
	}
	return nil
}
