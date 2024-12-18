package dao

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type SystemDao struct {
	db.BaseDao[model.SystemModel]
}

func NewSystemDao(ctx context.Context, dao ...*db.BaseDao[model.SystemModel]) *SystemDao {
	return &SystemDao{BaseDao: db.NewDao[model.SystemModel](ctx, dao...)}
}

func (dao *SystemDao) Create(system *model.SystemModel) error {
	// 检查是否存在相同Code的系统
	var existingSystem model.SystemModel
	err := dao.GetOrm().Where("code = ?", system.Code).First(&existingSystem).Error
	if err == nil {
		return errors.New("system code already exists")
	}
	
	return dao.GetOrm().Create(system).Error
}
