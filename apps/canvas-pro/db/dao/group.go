package dao

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

type GroupDao struct {
	db.BaseDao[model.GroupModel]
}

func NewGroupDao(ctx context.Context, dao ...*db.BaseDao[model.GroupModel]) *GroupDao {
	return &GroupDao{BaseDao: db.NewDao[model.GroupModel](ctx, dao...)}
}

func FindByNameAndWorkspace(tx *gorm.DB, name string, workspace uint64, groupTypeTag string) ([]*model.GroupModel, error) {

	orm := tx.Where("workspace = ? AND type_tag = ?", workspace, groupTypeTag)
	if name != "" {
		orm = orm.Where("name = ?", name)
	}
	res := []*model.GroupModel{}
	if err := orm.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

var ErrorNameExisted = errors.New("error group name is exist")

func Create(tx *gorm.DB, workspaceId uint64, name, typeTag string) (*model.GroupModel, error) {
	records, err := FindByNameAndWorkspace(tx, name, workspaceId, typeTag)
	if records != nil && err == nil {
		if len(records) > 0 {
			err = ErrorNameExisted
			return nil, err
		}
	}
	record := &model.GroupModel{}
	record.Name, record.Workspace, record.TypeTag = name, workspaceId, typeTag
	err = tx.Create(record).Error
	return record, err
}
