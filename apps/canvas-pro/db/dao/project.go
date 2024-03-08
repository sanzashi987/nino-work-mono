package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/apps/canvas-pro/enums"
	"github.com/cza14h/nino-work/pkg/db"
)

type ProjectDao struct {
	db.BaseDao[model.ProjectModel]
}

func NewProjectDao(ctx context.Context) *ProjectDao {
	return &ProjectDao{db.InitBaseDao[model.ProjectModel](ctx)}
}

func (dao *ProjectDao) GetList(page, size int, workspace string /**optional**/, name, group *string) (projects *[]model.ProjectModel, err error) {

	query := dao.DB.Scopes(db.Paginate(page, size)).Model(&model.ProjectModel{}).Where("workspace = ?", workspace)

	if group != nil {
		_, _, err = enums.GetIdFromCode(*group)
		if err != nil {
			return
		}

		query = query.Where(" group = ?", *group)
	}

	if name != nil {
		query = query.Where(" group LIKE ?", *name)
	}
	err = query.Find(projects).Error
	return
	// filterByName := func(db *gorm.DB) *gorm.DB {
	// 	res := db
	// 	if name != nil {
	// 		res = res.Where(" name LIKE ?", "%"+*name+"%")
	// 	}
	// 	return res
	// }
	// err = dao.DB.Preload("Projects", filterByName).Scopes(db.Paginate(page, size)).Model(&groupModel).Association("Projects").Find(projects)
	// return

}
