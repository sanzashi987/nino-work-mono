package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type DataSourceDao struct {
	db.BaseDao[model.DataSourceModel]
}

func NewDataSourceDao(ctx context.Context, dao ...*db.BaseDao[model.DataSourceModel]) *DataSourceDao {
	var baseDao db.BaseDao[model.DataSourceModel]
	if len(dao) > 0 {
		baseDao = *dao[0]
	} else {
		baseDao = db.InitBaseDao[model.DataSourceModel](ctx)
	}
	return &DataSourceDao{BaseDao: baseDao}
}

var dataSourceTableName = model.DataSourceModel{}.TableName()
