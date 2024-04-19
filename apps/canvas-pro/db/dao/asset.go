package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/consts"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type AssetDao struct {
	db.BaseDao[model.AssetModel]
}

func NewAssetDao(ctx context.Context, dao ...*db.BaseDao[model.AssetModel]) *AssetDao {
	var baseDao db.BaseDao[model.AssetModel]
	if len(dao) > 0 {
		baseDao = *dao[0]
	} else {
		baseDao = db.InitBaseDao[model.AssetModel](ctx)
	}
	return &AssetDao{BaseDao: baseDao}
}

var assetTableName = model.ProjectModel{}.TableName()

func (dao *AssetDao) DeleleGroupEffect(groupId, workspace uint64) error {
	return dao.GetOrm().Table(assetTableName).Where("group_id = ? AND workspace = ?", groupId, workspace).Updates(map[string]any{"group_id": 0}).Error
}

func (dao *AssetDao) BatchMoveGroup(groupId, workspace uint64, projectIds []uint64) error {
	orm := dao.GetOrm().Table(assetTableName)
	return orm.Where("id IN ? AND workspace = ?", projectIds, workspace).Update("group_id", groupId).Error
}

func (dao *AssetDao) CreateAsset(workspace uint64, fileId, assetType string) (err error) {
	toCreate := model.AssetModel{
		Version: consts.DefaultVersion,
		FileId:  fileId,
		Type:    assetType,
	}
	toCreate.Workspace, toCreate.TypeTag = workspace, assetType
	return dao.GetOrm().Table(assetTableName).Create(&toCreate).Error
}
