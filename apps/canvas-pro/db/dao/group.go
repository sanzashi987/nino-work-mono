package dao

import (
	"context"
	"errors"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type GroupDao struct {
	db.BaseDao[model.GroupModel]
}

func NewGroupDao(ctx context.Context, dao ...*db.BaseDao[model.GroupModel]) *GroupDao {
	var baseDao db.BaseDao[model.GroupModel]
	if len(dao) > 0 {
		baseDao = *dao[0]
	} else {
		baseDao = db.InitBaseDao[model.GroupModel](ctx)
	}
	return &GroupDao{BaseDao: baseDao}

	// return &GroupDao{db.InitBaseDao[model.GroupModel](ctx)}
}

func (dao *GroupDao) FindByNameAndWorkspace(name string, workspace uint64) (res []model.GroupModel, err error) {
	err = dao.GetOrm().Where("name = ? AND workspace = ?", name, workspace).Find(&res).Error
	return
}

func (dao *GroupDao) Delete(id uint64) (err error) {
	toDelete := model.GroupModel{}
	toDelete.Id = id

	if err = dao.GetOrm().Delete(&toDelete).Error; err != nil {
		return
	}
	return
}

var ErrorNameExisted = errors.New("error group name is exist")

func (dao *GroupDao) Create(workspaceId uint64, name, typeTag string) (record *model.GroupModel, err error) {
	records, err := dao.FindByNameAndWorkspace(name, workspaceId)
	if records != nil && err == nil {
		if len(records) > 0 {
			err = ErrorNameExisted
			return
		}
	}
	record = &model.GroupModel{}
	record.Name, record.Workspace, record.TypeTag = name, workspaceId, typeTag
	err = dao.GetOrm().Create(record).Error
	return
}
