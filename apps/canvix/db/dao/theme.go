package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type ThemeDao struct {
	db.BaseDao[model.ThemeModel]
}

func NewThemeDao(ctx context.Context, dao ...*db.BaseDao[model.ThemeModel]) *ThemeDao {
	return &ThemeDao{BaseDao: db.NewDao[model.ThemeModel](ctx, dao...)}
}

func (dao ThemeDao) GetWorkspaceThemes(workspaceId uint64) ([]model.ThemeModel, error) {
	res := []model.ThemeModel{}
	err := dao.GetOrm().Where("workspace = ? ", workspaceId).Find(&res).Error
	return res, err
}
