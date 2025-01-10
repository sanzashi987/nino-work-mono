package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/consts"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/model"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

type DataSourceService struct{}

var DataSourceServiceImpl *DataSourceService = &DataSourceService{}

type DataSourceDetail struct {
	SourceName string `json:"sourceName"`
	SourceType uint8  `json:"sourceType"`
	SourceInfo string `json:"sourceInfo"`
	SourceId   string `json:"sourceId"`
	Creator    string `json:"userIdentify"`
	shared.DBTime
}

func intoDataSourceDetail(input model.DataSourceModel) DataSourceDetail {
	return DataSourceDetail{
		SourceName: input.Name,
		SourceType: input.SourceType,
		SourceInfo: input.SourceInfo,
		SourceId:   strconv.FormatUint(input.Id, 10),
		Creator:    strconv.FormatUint(input.Creator, 10),
		DBTime: shared.DBTime{
			CreateTime: input.GetCreatedDate(),
			UpdateTime: input.GetUpdatedDate(),
		},
	}
}

func (serv *DataSourceService) ListDataSources(ctx context.Context, workspaceId uint64, page, size int, sourceName string, sourceType []string) ([]DataSourceDetail, error) {
	dataSourceDao := dao.NewDataSourceDao(ctx)

	records, err := dataSourceDao.FindByNameOrType(page, size, workspaceId, sourceName, sourceType)
	if err != nil {
		return nil, err
	}

	response := []DataSourceDetail{}
	for _, source := range records {
		temp := intoDataSourceDetail(source)

		response = append(response, temp)
	}

	return response, nil
}

func (serv *DataSourceService) GetDataSourceById(ctx context.Context, workspaceId uint64, sourceIdCode string) (res DataSourceDetail, err error) {
	dataSourceDao := dao.NewDataSourceDao(ctx)

	var id uint64
	if id, _, err = consts.GetIdFromCode(sourceIdCode); err != nil {
		return
	}

	record, err := dataSourceDao.GetDataSourceById(id)
	if err != nil {
		return
	}
	res = intoDataSourceDetail(record)

	return
}

type UpdateDataSourceRequest struct {
	SourceName string `json:"sourceName"`
	SourceInfo string `json:"sourceInfo"`
	SourceId   string `json:"sourceId" binding:"required"`
}

func (serv *DataSourceService) Update(ctx context.Context, workspaceId uint64, payload *UpdateDataSourceRequest) (res DataSourceDetail, err error) {
	dataSourceDao := dao.NewDataSourceDao(ctx)

	var id uint64
	if id, _, err = consts.GetIdFromCode(payload.SourceId); err != nil {
		return
	}

	if payload.SourceName == "" && payload.SourceInfo == "" {
		err = errors.New("sourceName and sourceInfo are both empty")
		return
	}

	record, err := dataSourceDao.Update(workspaceId, id, payload.SourceName, payload.SourceInfo)
	if err != nil {
		return
	}
	res = intoDataSourceDetail(record)

	return
}

type CreateDataSourceRequest struct {
	SourceName string `json:"sourceName" binding:"required"`
	SourceType string `json:"sourceType" binding:"required"`
	SourceInfo string `json:"sourceInfo" binding:"required"`
}

func (serv *DataSourceService) Create(ctx context.Context, workspaceId uint64, payload *CreateDataSourceRequest) (res DataSourceDetail, err error) {
	dataSourceDao := dao.NewDataSourceDao(ctx)
	sourceTypeEnum, exist := model.SourceTypeStringToEnum[payload.SourceType]
	if !exist {
		err = errors.New("source type not found")
		return
	}

	record, err := dataSourceDao.Create(workspaceId, sourceTypeEnum, payload.SourceName, payload.SourceInfo)
	if err != nil {
		return
	}
	res = intoDataSourceDetail(record)

	return
}

func (serv *DataSourceService) Delete(ctx context.Context, workspaceId uint64, codes []string) error {
	dataSourceDao := dao.NewDataSourceDao(ctx)

	ids := []uint64{}
	for _, code := range codes {
		id, tag, err := consts.GetIdFromCode(code)
		if err != nil {
			return err
		}
		if tag != consts.DATASOURCE {
			return errors.New("Has invalid datasource code: " + code)
		}
		ids = append(ids, id)
	}

	return dataSourceDao.Delete(workspaceId, ids)
}
