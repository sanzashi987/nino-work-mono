package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type ThemeDao struct {
	db.BaseDao[model.ThemeModel]
}

func NewThemeDao(ctx context.Context, dao ...*db.BaseDao[model.ThemeModel]) *ThemeDao {
	return &ThemeDao{BaseDao: db.NewDao[model.ThemeModel](ctx, dao...)}
}

var themeTableName = model.ThemeModel{}.TableName()

func (dao ThemeDao) BatchDeleleTheme(workspaceId uint64, ids []uint64) error {
	return dao.GetOrm().Table(themeTableName).Where("id in ? AND workspace = ?", ids, workspaceId).Delete(&model.ThemeModel{}).Error
}

func (dao ThemeDao) CreateUserTheme(workspaceId uint64, name, config string) error {

	toCreate := model.ThemeModel{
		Type:   consts.CustomizedTheme,
		Config: config,
	}
	toCreate.Workspace = workspaceId
	toCreate.Name, toCreate.TypeTag = name, consts.THEME

	return dao.GetOrm().Create(&toCreate).Error
}

func (dao ThemeDao) UpdateUserTheme(workspaceId, id uint64, name, config *string) error {

	toUpdate := map[string]string{}

	if name != nil {
		toUpdate["name"] = *name
	}

	if config != nil {
		toUpdate["config"] = *config
	}

	return dao.GetOrm().Model(&model.ThemeModel{}).Where("id = ? and workspace = ?", id, workspaceId).Updates(toUpdate).Error
}

func (dao ThemeDao) GetWorkspaceThemes(workspaceId uint64) ([]model.ThemeModel, error) {
	res := []model.ThemeModel{}
	err := dao.GetOrm().Where("workspace = ? ", workspaceId).Find(&res).Error
	return res, err
}
