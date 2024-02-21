package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type TemplateDao struct {
	db.BaseDao[model.TemplateModel]
}

func NewTemplateDao(ctx context.Context) *TemplateDao {
	return &TemplateDao{db.InitBaseDao[model.TemplateModel](ctx)}
}
