package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type TemplateDao struct {
	db.BaseDao[model.TemplateModel]
}

func NewTemplateDao(ctx context.Context, dao ...*db.BaseDao[model.TemplateModel]) *TemplateDao {
	return &TemplateDao{BaseDao: db.NewDao[model.TemplateModel](ctx, dao...)}
}
