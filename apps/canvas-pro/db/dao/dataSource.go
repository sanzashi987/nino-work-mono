package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/consts"
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

func (serv *DataSourceDao) FindByNameOrType(page, size int, workspace uint64, name string, sourceType []string) (result []model.DataSourceModel, err error) {
	query := serv.GetOrm().Scopes(db.Paginate(page, size)).Model(&model.DataSourceModel{}).Where("workspace = ?", workspace)

	if name != "" {
		query = query.Where("name LIKE ?", name)
	}

	if len(sourceType) != 0 {
		sourceTypeEnums := []uint8{}

		for _, v := range sourceType {
			enum, exist := model.SourceTypeStringToEnum[v]
			if exist {
				sourceTypeEnums = append(sourceTypeEnums, enum)
			}
		}

		query = query.Where("source_type IN ?", sourceTypeEnums)
	}

	err = query.Find(&result).Error
	return

}

func (serv *DataSourceDao) GetDataSourceById(id uint64) (model.DataSourceModel, error) {
	var result model.DataSourceModel
	err := serv.GetOrm().Model(&model.DataSourceModel{}).Where("id = ?", id).First(&result).Error
	return result, err
}

func (serv *DataSourceDao) Update(workspace, id uint64, sourceName, sourceInfo string) (model.DataSourceModel, error) {
	toUpdate := map[string]string{}
	var result = model.DataSourceModel{}
	if sourceName != "" {
		toUpdate["name"] = sourceName
	}

	if sourceInfo != "" {
		toUpdate["source_info"] = sourceInfo
	}
	err := serv.GetOrm().Model(&result).Where("id = ? and workspace = ?", id, workspace).Updates(toUpdate).Error

	return result, err
}

func (serv *DataSourceDao) Create(workspace uint64, sourceType uint8, name, info string) (model.DataSourceModel, error) {
	toCreate := model.DataSourceModel{
		Version:    consts.DefaultVersion,
		SourceType: sourceType,
		SourceInfo: info,
	}
	toCreate.Workspace, toCreate.TypeTag = workspace, consts.DATASOURCE
	err := serv.GetOrm().Table(dataSourceTableName).Create(&toCreate).Error

	return toCreate, err
}

// func (serv * DataSourceDao) Delete