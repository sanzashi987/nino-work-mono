package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type ThemeDao struct {
	db.BaseDao[model.ThemeModel]
}

func NewThemeDao(ctx context.Context) *ThemeDao {
	return &ThemeDao{
		db.InitBaseDao[model.ThemeModel](ctx),
	}
}
