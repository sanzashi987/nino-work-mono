package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type PermissionDao struct {
	db.BaseDao[model.PermissionModel]
}

func NewPermissionDao(ctx context.Context, dao ...*db.BaseDao[model.PermissionModel]) *PermissionDao {
	return &PermissionDao{BaseDao: db.NewDao[model.PermissionModel](ctx, dao...)}
}

func (dao *PermissionDao) CreatePermission(permission *model.PermissionModel) error {
	return dao.GetOrm().Create(permission).Error
}
