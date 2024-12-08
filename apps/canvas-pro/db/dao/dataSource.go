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

func (serv *DataSourceDao) FindByNameOrType(page, size int, workspace uint64, name, sourceType string) (result []model.DataSourceModel, err error) {
	query := serv.GetOrm().Scopes(db.Paginate(page, size)).Model(&model.DataSourceModel{}).Where("workspace = ?", workspace)

	if name != "" {
		query = query.Where("name LIKE ?", name)
	}

	if sourceType != "" {
		query = query.Where("source_type = ?", sourceType)
	}

	err = query.Find(&result).Error
	return

}

func (serv *DataSourceDao) GetDataSourceById(id uint64) (model.DataSourceModel, error) {
	var result model.DataSourceModel
	err := serv.GetOrm().Model(&model.DataSourceModel{}).Where("id = ?", id).First(&result).Error
	return result, err
}
