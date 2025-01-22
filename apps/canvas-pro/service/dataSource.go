package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/consts"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
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

func intoDataSourceDetail(input *model.DataSourceModel) *DataSourceDetail {
	return &DataSourceDetail{
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

type QueryDataSourceSearchRequest struct {
	SourceName string   `json:"sourceName"`
	SourceType []string `json:"sourceType"`
	Search     string   `json:"search"`
}

type QueryDataSourceRequest struct {
	shared.PaginationRequest
	QueryDataSourceSearchRequest
}

func (serv *DataSourceService) ListDataSources(ctx context.Context, workspaceId uint64, payload *QueryDataSourceRequest) ([]*DataSourceDetail, error) {
	tx := db.NewTx(ctx)

	records, err := dao.FindByNameOrType(tx, workspaceId, payload.SourceName, payload.SourceType, payload.PaginationRequest)
	if err != nil {
		return nil, err
	}

	response := []*DataSourceDetail{}
	for _, source := range records {
		temp := intoDataSourceDetail(source)

		response = append(response, temp)
	}

	return response, nil
}

func (serv *DataSourceService) GetDataSourceById(ctx context.Context, workspaceId uint64, sourceIdCode string) (*DataSourceDetail, error) {
	tx := db.NewTx(ctx)

	id, _, err := consts.GetIdFromCode(sourceIdCode)
	if err != nil {
		return nil, err
	}

	result := model.DataSourceModel{}

	if err := tx.Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return intoDataSourceDetail(&result), nil
}

type UpdateDataSourceRequest struct {
	SourceName *string `json:"sourceName"`
	SourceInfo *string `json:"sourceInfo"`
	SourceId   string  `json:"sourceId" binding:"required"`
}

func (serv *DataSourceService) Update(ctx context.Context, workspaceId uint64, payload *UpdateDataSourceRequest) (*DataSourceDetail, error) {
	tx := db.NewTx(ctx)

	id, _, err := consts.GetIdFromCode(payload.SourceId)
	if err != nil {
		return nil, err
	}

	if payload.SourceName == nil && payload.SourceInfo == nil {
		return nil, errors.New("sourceName and sourceInfo are both empty")
	}

	record, err := dao.UpdateProject(tx, workspaceId, id, payload.SourceName, payload.SourceInfo)
	if err != nil {
		return nil, err
	}

	return intoDataSourceDetail(record), nil
}

type CreateDataSourceRequest struct {
	SourceName string `json:"sourceName" binding:"required"`
	SourceType string `json:"sourceType" binding:"required"`
	SourceInfo string `json:"sourceInfo" binding:"required"`
}

func (serv *DataSourceService) Create(ctx context.Context, workspaceId uint64, payload *CreateDataSourceRequest) (*DataSourceDetail, error) {
	tx := db.NewTx(ctx)

	sourceTypeEnum, exist := model.SourceTypeStringToEnum[payload.SourceType]
	if !exist {
		return nil, errors.New("source type not found")
	}

	toCreate := model.DataSourceModel{
		Version:    consts.DefaultVersion,
		SourceType: sourceTypeEnum,
		SourceInfo: payload.SourceInfo,
	}
	toCreate.Workspace, toCreate.TypeTag, toCreate.Name = workspaceId, consts.DATASOURCE, payload.SourceName

	if err := tx.Create(&toCreate).Error; err != nil {
		return nil, err
	}

	return intoDataSourceDetail(&toCreate), nil
}

func (serv *DataSourceService) Delete(ctx context.Context, workspaceId uint64, codes []string) error {
	tx := db.NewTx(ctx)

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

	return tx.Model(&model.DataSourceModel{}).Where("workspace = ? and id in ?", workspaceId, ids).Delete(&model.DataSourceModel{}).Error
}
