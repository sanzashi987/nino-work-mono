package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type TemplateDao struct {
	db.BaseDao[model.TemplateModel]
}

func NewTemplateDao(ctx context.Context) *TemplateDao {
	return &TemplateDao{db.InitBaseDao[model.TemplateModel](ctx)}
}
