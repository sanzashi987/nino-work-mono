package service

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type ThemeService struct{}

var ThemeServiceImpl *ThemeService = &ThemeService{}

func (serv ThemeService) UpdateTheme(ctx context.Context, workspaceId, themeId uint64, name, config *string) error {
	themeDao := dao.NewThemeDao(ctx)
	return themeDao.UpdateUserTheme(workspaceId, themeId, name, config)
}

type CreateThemeReq struct {
	Name   string `json:"name" binding:"required"`
	Config string `json:"theme" binding:"required"`
	/**
	 * 0: default theme
	 * 1: custom theme
	 */
	Flag int `json:"flag" binding:"required"`
}

func CreateTheme(ctx context.Context, workspaceId uint64, req *CreateThemeReq) error {
	tx := db.NewTx(ctx)

	toCreate := model.ThemeModel{
		Type:   int8(req.Flag),
		Config: req.Config,
	}

	toCreate.Workspace = workspaceId
	toCreate.Name, toCreate.TypeTag = req.Name, consts.THEME

	tx = tx.Begin()

	if err := tx.Model(&model.ThemeModel{}).Create(&toCreate).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (serv ThemeService) GetThemes(ctx context.Context, workspaceId uint64) ([]model.ThemeModel, error) {
	themeDao := dao.NewThemeDao(ctx)
	return themeDao.GetWorkspaceThemes(workspaceId)
}

func (serv ThemeService) DeleteThemes(ctx context.Context, workspaceId uint64, ids []uint64) error {
	themeDao := dao.NewThemeDao(ctx)
	return themeDao.BatchDeleleTheme(workspaceId, ids)
}
