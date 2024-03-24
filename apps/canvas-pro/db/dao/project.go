package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/consts"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type ProjectDao struct {
	db.BaseDao[model.ProjectModel]
}

func NewProjectDao(ctx context.Context) *ProjectDao {
	return &ProjectDao{db.InitBaseDao[model.ProjectModel](ctx)}
}

func (dao *ProjectDao) GetList(page, size int, workspace string /**optional**/, name, group *string) (projects *[]model.ProjectModel, err error) {

	query := dao.GetOrm().Scopes(db.Paginate(page, size)).Model(&model.ProjectModel{}).Where("workspace = ?", workspace)

	if group != nil {
		_, _, err = consts.GetIdFromCode(*group)
		if err != nil {
			return
		}

		query = query.Where(" group_id = ?", *group)
	}

	if name != nil {
		query = query.Where(" name LIKE ?", *name)
	}
	err = query.Find(projects).Error
	return
}

func (dao *ProjectDao) BatchLogicalDelete(ids []uint64) error {
	return dao.GetOrm().Table(model.ProjectModel{}.TableName()).Where("id IN ?", ids).Update("deleted", db.Deleted).Error
}
