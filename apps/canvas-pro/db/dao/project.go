package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
	"gorm.io/gorm"
)

type ProjectDao struct {
	db.BaseDao[model.ProjectModel]
}

func NewProjectDao(ctx context.Context) *ProjectDao {
	return &ProjectDao{db.InitBaseDao[model.ProjectModel](ctx)}
}

func (p *ProjectDao) GetList(page, size int, name, group, workspace string) (projects *[]model.ProjectModel, err error) {

	groupModel := model.ProjectGroupModel{
		BaseModel: model.BaseModel{
			Workspace: workspace,
		},
	}

	filterByName := func(db *gorm.DB) *gorm.DB {
		res := db
		if name != "" {
			res = res.Where(" name LIKE ?", "%"+name+"%")
		}
		return res
	}
	err = p.DB.Preload("Projects", filterByName).Scopes(db.Paginate(page, size)).Model(&groupModel).Association("Projects").Find(projects)
	return
}
