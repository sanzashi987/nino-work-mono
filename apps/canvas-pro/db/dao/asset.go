package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/consts"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
	"gorm.io/gorm"
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

func (dao *AssetDao) CreateAsset(workspace, groupId uint64, name, fileId, assetType string) (*model.AssetModel, error) {
	toCreate := model.AssetModel{
		Version: consts.DefaultVersion,
		FileId:  fileId,
		GroupId: groupId,
	}
	toCreate.Workspace, toCreate.TypeTag = workspace, assetType
	err := dao.GetOrm().Table(assetTableName).Create(&toCreate).Error

	return &toCreate, err
}

type GroupCount struct {
	Id    uint64
	Count uint64
}

func (dao AssetDao) GetAssetCountByGroup(workspaceId uint64, groupIds []uint64) (res []GroupCount, err error) {

	err = dao.GetOrm().Table(assetTableName).Where("workspace = ?", workspaceId).Where("group_id IN ?", groupIds).Select("id", "COUNT(id) as count").Group("group_id").Find(&res).Error
	return

}

func (dao AssetDao) update(workspaceId, assetId uint64) *gorm.DB {
	return dao.GetOrm().Table(assetTableName).Where("id = ?", assetId).Where("workspace = ?", workspaceId)

}

func (dao AssetDao) UpdateAssetName(workspaceId, assetId uint64, assetName string) error {
	return dao.update(workspaceId, assetId).Update("name", assetName).Error
}

func (dao AssetDao) ListAssets(workspaceId uint64, groupId *uint64, page, size int, typeTag string) ([]model.AssetModel, error) {
	res := []model.AssetModel{}
	orm := dao.GetOrm().Scopes(db.Paginate(page, size)).Table(assetTableName).Where("workspace = ? ", workspaceId).Where("type_tag = ?", typeTag)
	if groupId != nil {
		orm = orm.Where("group_id = ? ", groupId)
	}

	err := orm.Find(&res).Error
	return res, err
}

func (dao AssetDao) GetAssetCount(workspaceId uint64, groupId *uint64, page, size int, typeTag string) (int64, error) {

	orm := dao.GetOrm().Table(assetTableName).Select("id").Where("workspace = ? ", workspaceId).Where("type_tag = ?", typeTag)
	if groupId != nil {
		orm = orm.Where("group_id = ? ", groupId)
	}
	var recordCount *int64
	err := orm.Count(recordCount).Error
	return *recordCount, err
}

func (dao AssetDao) GetSingleAsset(workspaceId, assetId uint64) (res model.AssetModel, err error) {
	err = dao.GetOrm().Table(assetTableName).Where("id = ?", assetId).Where("workspace = ?", workspaceId).Take(&res).Error
	return
}
