package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type MenuDao struct {
	db.BaseDao[model.MenuModel]
}

func NewMenuDao(ctx context.Context, dao ...*db.BaseDao[model.MenuModel]) *MenuDao {
	return &MenuDao{BaseDao: db.NewDao(ctx, dao...)}
}

func (dao *MenuDao) GetMenusByRoles(roles *[]model.RoleModel) error {
	return dao.GetOrm().Preload("Menus").Find(roles).Error
}
