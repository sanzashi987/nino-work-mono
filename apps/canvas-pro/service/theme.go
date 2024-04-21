package service

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
)

type ThemeService struct{}

var ThemeServiceImpl *ThemeService = &ThemeService{}

func (serv ThemeService) UpdateTheme(ctx context.Context, workspaceId, themeId uint64, name, config *string) error {
	themeDao := dao.NewThemeDao(ctx)
	return themeDao.UpdateUserTheme(workspaceId, themeId, name, config)
}
