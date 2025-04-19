package dao

import (
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type ThemeDao struct {
	db.BaseDao[model.ThemeModel]
}

func (dao ThemeDao) GetWorkspaceThemes(workspaceId uint64) ([]model.ThemeModel, error) {
	res := []model.ThemeModel{}
	err := dao.GetOrm().Where("workspace = ? ", workspaceId).Find(&res).Error
	return res, err
}
