package dao

import (
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"gorm.io/gorm"
)

func UpdateProject(tx *gorm.DB, workspace, id uint64, sourceName, sourceInfo *string) (*model.DataSourceModel, error) {
	toUpdate := map[string]string{}
	var result = model.DataSourceModel{}
	if sourceName != nil {
		toUpdate["name"] = *sourceName
	}

	if sourceInfo != nil {
		toUpdate["source_info"] = *sourceInfo
	}
	err := tx.Model(&result).Where("id = ? and workspace = ?", id, workspace).Updates(toUpdate).Error

	return &result, err
}
