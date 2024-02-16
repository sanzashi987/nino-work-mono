package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type AssetDao struct {
	db.BaseDao[model.AssetModel]
}

func NewAssetDao(ctx context.Context) *AssetDao {
	return &AssetDao{db.InitBaseDao[model.AssetModel](ctx)}
}
