package dao

import (
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"gorm.io/gorm"
)

var projectTableName = model.ProjectModel{}.TableName()

func BatchLogicalDelete(tx *gorm.DB, ids []uint64) error {
	return tx.Table(projectTableName).Where("id IN ?", ids).Delete(&model.ProjectModel{}).Error

}

func ProjectDeleleGroupEffect(tx *gorm.DB, groupId, workspace uint64) error {
	return tx.Table(projectTableName).Where("group_id = ? AND workspace = ?", groupId, workspace).Updates(map[string]any{"group_id": 0}).Error
}
