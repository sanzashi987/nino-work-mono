package dao

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type ApplicationDao struct {
	db.BaseDao[model.ApplicationModel]
}

func NewApplicationDao(ctx context.Context, dao ...*db.BaseDao[model.ApplicationModel]) *ApplicationDao {
	return &ApplicationDao{BaseDao: db.NewDao(ctx, dao...)}
}

func (dao *ApplicationDao) Create(system *model.ApplicationModel) error {
	// 检查是否存在相同Code的系统
	var existingSystem model.ApplicationModel
	err := dao.GetOrm().Where("code = ?", system.Code).First(&existingSystem).Error
	if err == nil {
		return errors.New("system code already exists")
	}

	return dao.GetOrm().Create(system).Error
}

func (dao *ApplicationDao) InitPermissionForSystem(app *model.ApplicationModel, super *model.PermissionModel, admin *model.PermissionModel) error {

	toUpdate := map[string]any{
		"super_admin": super.Id,
		"authorizer":  admin.Id,
	}
	err := dao.GetOrm().Model(app).Updates(toUpdate).Error
	return err
}
