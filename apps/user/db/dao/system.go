package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type SystemDao struct {
	db.BaseDao[model.SystemModel]
}

func NewSystemDao(ctx context.Context, dao ...*db.BaseDao[model.SystemModel]) *SystemDao {
	return &SystemDao{BaseDao: db.NewDao[model.SystemModel](ctx, dao...)}
}

func (dao *SystemDao) CreatePermission(permission *model.PermissionModel) error {
	return dao.GetOrm().Create(permission).Error
}
