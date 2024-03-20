package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/consts"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type GroupDao struct {
	db.BaseDao[model.BaseModel]
}

func NewGroupDao(ctx context.Context) *GroupDao {
	return &GroupDao{db.InitBaseDao[model.BaseModel](ctx)}
}

type DBModel interface {
	TableName() string
}

func (dao *GroupDao) FindByNameAndWorkspace(name, workspace string) (res *[]model.BaseModel, err error) {
	err = dao.DB.Where("name = ? AND workspace = ?", name, workspace).Find(res).Error
	return
}

func (dao *GroupDao) Create(name, workspace string, dbModel DBModel) error {
	newGroup := model.BaseModel{Name: name, Workspace: workspace, TypeTag: consts.GROUP}
	return dao.DB.Table(dbModel.TableName()).Create(&newGroup).Error
}

func (dao *GroupDao) Delete(name, workspace string, dbModel DBModel) error {
	newGroup := model.BaseModel{Name: name, Workspace: workspace, Deleted: model.Deleted}
	return dao.DB.Table(dbModel.TableName()).Create(&newGroup).Error
}
