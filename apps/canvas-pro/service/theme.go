package service

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
)

type ThemeService struct{}

var ThemeServiceImpl *ThemeService = &ThemeService{}

func (serv ThemeService) UpdateTheme(ctx context.Context, workspaceId, themeId uint64, name, config *string) error {
	themeDao := dao.NewThemeDao(ctx)
	return themeDao.UpdateUserTheme(workspaceId, themeId, name, config)
}

func (serv ThemeService) CreateTheme(ctx context.Context, workspaceId uint64, name, config string) error {
	themeDao := dao.NewThemeDao(ctx)
	return themeDao.CreateUserTheme(workspaceId, name, config)
}

func (serv ThemeService) GetThemes(ctx context.Context, workspaceId uint64) ([]model.ThemeModel, error) {
	themeDao := dao.NewThemeDao(ctx)
	return themeDao.GetWorkspaceThemes(workspaceId)
}

func (serv ThemeService) DeleteThemes(ctx context.Context, workspaceId uint64, ids []uint64) error {
	themeDao := dao.NewThemeDao(ctx)
	return themeDao.BatchDeleleTheme(workspaceId, ids)
}
