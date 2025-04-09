package dao

import (
	"errors"

	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"gorm.io/gorm"
)

func FindByNameAndWorkspace(tx *gorm.DB, workspace uint64, name, typeTag string) ([]*model.GroupModel, error) {

	orm := tx.Where("workspace = ? AND type_tag = ?", workspace, typeTag)
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
	records, err := FindByNameAndWorkspace(tx, workspaceId, name, typeTag)
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
