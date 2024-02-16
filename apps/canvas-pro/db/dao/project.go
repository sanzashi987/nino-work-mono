package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type ProjectDao struct {
	db.BaseDao[model.ProjectModel]
}

func NewProjectDao(ctx context.Context) *ProjectDao {
	return &ProjectDao{db.InitBaseDao[model.ProjectModel](ctx)}
}

func (p *ProjectDao) Create(project *model.ProjectModel) {

}
