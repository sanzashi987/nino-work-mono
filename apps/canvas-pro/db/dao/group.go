package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/apps/canvas-pro/enums"
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

func (dao *GroupDao) Create(workspace, name string, dbModel DBModel) error {
	newGroup := model.BaseModel{Name: name, Workspace: workspace, TypeTag: enums.GROUP}
	return dao.DB.Table(dbModel.TableName()).Create(&newGroup).Error
}
