package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type AssetDao struct {
	db.BaseDao[model.AssetModel]
}

func NewAssetDao(ctx context.Context, dao ...db.BaseDao[model.AssetModel]) *AssetDao {
	var baseDao db.BaseDao[model.AssetModel]
	if len(dao) > 0 {
		baseDao = dao[0]
	} else {
		baseDao = db.InitBaseDao[model.AssetModel](ctx)
	}
	return &AssetDao{BaseDao: baseDao}
}

func (dao *AssetDao) DeleleGroupEffect(groupId, workspace uint64) {
	orm := dao.GetOrm()
	selector := model.AssetModel{}
	selector.GroupId, selector.Workspace = groupId, workspace
	// res := &[]model.AssetModel{}
	orm.Model(selector).Update("group_id", 0)

}
