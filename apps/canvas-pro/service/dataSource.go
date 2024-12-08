package service

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/http/request"
)

type DataSourceService struct{}

var DataSourceServiceImpl *DataSourceService = &DataSourceService{}

// data source

type ListDataSourcesResponse struct {
	SourceName string `json:"sourceName"`
	SourceType string `json:"sourceType"`
	SourceInfo string `json:"sourceInfo"`
	SourceId   string `json:"sourceId"`
	Createor   string `json:"userIdentify"`
	request.DBTime
}

func (serv DataSourceService) ListDataSources(ctx context.Context, workspaceId uint64, page, size int, sourceName, sourceType string) ([]ListDataSourcesResponse, error) {
	dataSourceDao := dao.NewDataSourceDao(ctx)

}
