package dao

import (
	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

func GetList(tx *gorm.DB, page, size int, workspace uint64, name *string, groupId *uint64) (projects *[]model.ProjectModel, err error) {

	query := tx.Scopes(db.Paginate(page, size)).Model(&model.ProjectModel{}).Where("workspace = ?", workspace)

	if groupId != nil {
		query = query.Where(" group_id = ?", *groupId)
	}

	if name != nil {
		query = query.Where(" name LIKE ?", *name)
	}
	err = query.Find(projects).Error
	return
}

var projectTableName = model.ProjectModel{}.TableName()

func BatchLogicalDelete(tx *gorm.DB, ids []uint64) error {
	return tx.Table(projectTableName).Where("id IN ?", ids).Delete(&model.ProjectModel{}).Error

}

func ProjectDeleleGroupEffect(tx *gorm.DB, groupId, workspace uint64) error {
	return tx.Table(projectTableName).Where("group_id = ? AND workspace = ?", groupId, workspace).Updates(map[string]any{"group_id": 0}).Error
}
 