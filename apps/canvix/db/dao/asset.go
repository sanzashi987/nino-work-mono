package dao

import (
	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
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

func ListAssets(tx *gorm.DB, workspaceId uint64, groupId *uint64, page, size int, typeTag string) ([]model.AssetModel, error) {
	res := []model.AssetModel{}
	orm := tx.Model(&model.AssetModel{}).Scopes(db.Paginate(page, size)).Where("workspace = ? ", workspaceId).Where("type_tag = ?", typeTag)
	if groupId != nil {
		orm = orm.Where("group_id = ? ", groupId)
	}

	err := orm.Find(&res).Error
	return res, err
}

func GetAssetCount(tx *gorm.DB, workspaceId uint64, groupId *uint64, typeTag string) (int64, error) {

	orm := tx.Model(&model.AssetModel{}).Select("id").Where("workspace = ? ", workspaceId).Where("type_tag = ?", typeTag)
	if groupId != nil {
		orm = orm.Where("group_id = ? ", groupId)
	}
	var recordCount *int64
	err := orm.Count(recordCount).Error
	return *recordCount, err
}
