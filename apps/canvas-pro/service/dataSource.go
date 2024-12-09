package service

import (
	"context"
	"strconv"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/consts"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/model"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/http/request"
)

type DataSourceService struct{}

var DataSourceServiceImpl *DataSourceService = &DataSourceService{}

type DataSourceDetail struct {
	SourceName string `json:"sourceName"`
	SourceType uint8  `json:"sourceType"`
	SourceInfo string `json:"sourceInfo"`
	SourceId   string `json:"sourceId"`
	Creator    string `json:"userIdentify"`
	request.DBTime
}

func intoDataSourceDetail(input model.DataSourceModel) DataSourceDetail {
	return DataSourceDetail{
		SourceName: input.Name,
		SourceType: input.SourceType,
		SourceInfo: input.SourceConfig,
		SourceId:   strconv.FormatUint(input.Id, 10),
		Creator:    strconv.FormatUint(input.Creator, 10),
		DBTime: request.DBTime{
			CreateTime: input.GetCreatedDate(),
			UpdateTime: input.GetUpdatedDate(),
		},
	}
}

func (serv *DataSourceService) ListDataSources(ctx context.Context, workspaceId uint64, page, size int, sourceName string, sourceType []string) ([]DataSourceDetail, error) {
	dataSourceDao := dao.NewDataSourceDao(ctx)

	dataSources, err := dataSourceDao.FindByNameOrType(page, size, workspaceId, sourceName, sourceType)
	if err != nil {
		return nil, err
	}

	response := []DataSourceDetail{}
	for _, source := range dataSources {
		temp := intoDataSourceDetail(source)

		response = append(response, temp)
	}

	return response, nil
}

func (serv *DataSourceService) GetDataSourceById(ctx context.Context, workspaceId uint64, sourceIdCode string) (res *DataSourceDetail, err error) {
	dataSourceDao := dao.NewDataSourceDao(ctx)

	var id uint64
	if id, _, err = consts.GetIdFromCode(sourceIdCode); err != nil {
		return
	}

	dataSource, err := dataSourceDao.GetDataSourceById(id)
	if err != nil {
		return
	}
	temp := intoDataSourceDetail(dataSource)
	res = &temp

	return
}

type UpdateDataSourceRequest struct {
	SourceName string `json:"sourceName"`
	SourceType string `json:"sourceType"`
	SourceInfo string `json:"sourceInfo"`
	SourceId   string `json:"sourceId" binding:"required"`
}

func (serv *DataSourceService) UpdateDataSourceById(ctx context.Context, workspaceId uint64, payload *UpdateDataSourceRequest) (err error) {
	dataSourceDao := dao.NewDataSourceDao(ctx)

	var id uint64
	if id, _, err = consts.GetIdFromCode(payload.SourceId); err != nil {
		return
	}

	err = dataSourceDao.UpdateDataSourceById(workspaceId, id, payload.SourceName, payload.SourceType, payload.SourceInfo)
	if err != nil {
		return
	}

	return
}
