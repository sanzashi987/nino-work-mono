package dao

import (
	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"gorm.io/gorm"
)

func AssetDeleleGroupEffect(tx *gorm.DB, groupId, workspace uint64) error {
	return tx.Model(&model.AssetModel{}).Where("group_id = ? AND workspace = ?", groupId, workspace).Updates(map[string]any{"group_id": 0}).Error
}

func CreateAsset(tx *gorm.DB, workspace, groupId uint64, name, fileId, assetType string) (*model.AssetModel, error) {
	toCreate := model.AssetModel{
		Version: consts.DefaultVersion,
		FileId:  fileId,
		GroupId: groupId,
	}
	toCreate.Workspace, toCreate.TypeTag = workspace, assetType
	err := tx.Model(&model.AssetModel{}).Create(&toCreate).Error

	return &toCreate, err
}

func update(tx *gorm.DB, workspaceId, assetId uint64) *gorm.DB {
	return tx.Model(&model.AssetModel{}).Where("id = ?", assetId).Where("workspace = ?", workspaceId)

}

func UpdateAssetName(tx *gorm.DB, workspaceId, assetId uint64, assetName string) error {
	return update(tx, workspaceId, assetId).Update("name", assetName).Error
}
