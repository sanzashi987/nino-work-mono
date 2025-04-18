package dao

import (
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

func GetList(tx *gorm.DB, page, size int, workspace uint64, name *string, groupId *uint64) ([]*model.ProjectModel, int, error) {
	singlePage := tx.Scopes(db.Paginate(page, size)).Model(&model.ProjectModel{})
	allCount := tx.Model(&model.ProjectModel{})

	singlePage = singlePage.Where("workspace = ?", workspace)
	allCount = allCount.Where("workspace = ?", workspace)

	if groupId != nil {
		singlePage = singlePage.Where(" group_id = ?", *groupId)
		allCount = allCount.Where(" group_id = ?", *groupId)
	}

	if name != nil {
		singlePage = singlePage.Where(" name LIKE ?", *name)
		allCount = allCount.Where(" name LIKE ?", *name)
	}

	var totalCount int64
	if err := allCount.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	projects := []*model.ProjectModel{}
	if err := singlePage.Find(&projects).Error; err != nil {
		return nil, 0, err
	}

	return projects, int(totalCount), nil
}

var projectTableName = model.ProjectModel{}.TableName()

func BatchLogicalDelete(tx *gorm.DB, ids []uint64) error {
	return tx.Table(projectTableName).Where("id IN ?", ids).Delete(&model.ProjectModel{}).Error

}

func ProjectDeleleGroupEffect(tx *gorm.DB, groupId, workspace uint64) error {
	return tx.Table(projectTableName).Where("group_id = ? AND workspace = ?", groupId, workspace).Updates(map[string]any{"group_id": 0}).Error
}
