package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type GroupDao struct {
	db.BaseDao[model.GroupModel]
}

func NewGroupDao(ctx context.Context) *GroupDao {
	return &GroupDao{db.InitBaseDao[model.GroupModel](ctx)}
}

func (dao *GroupDao) FindByNameAndWorkspace(name, workspace string) (res *[]model.GroupModel, err error) {
	err = dao.GetOrm().Where("name = ? AND workspace = ?", name, workspace).Find(res).Error
	return
}

func (dao *GroupDao) Delete(id uint64, typeTag string) (err error) {
	toDelete := model.GroupModel{TypeTag: typeTag}
	toDelete.Id = id

	if err = dao.LogicalDelete(toDelete); err != nil {
		return
	}
	return
}
