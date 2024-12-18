package dao

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type GroupDao struct {
	db.BaseDao[model.GroupModel]
}

func NewGroupDao(ctx context.Context, dao ...*db.BaseDao[model.GroupModel]) *GroupDao {
	return &GroupDao{BaseDao: db.NewDao[model.GroupModel](ctx, dao...)}
}

func (dao *GroupDao) FindByNameAndWorkspace(name string, workspace uint64, groupTypeTag string) (res []model.GroupModel, err error) {

	orm := dao.GetOrm().Where(" workspace = ? AND type_tag = ?", workspace, groupTypeTag)
	if name != "" {
		orm = orm.Where("name = ?", name)
	}

	err = orm.Find(&res).Error
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
	records, err := dao.FindByNameAndWorkspace(name, workspaceId, typeTag)
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
