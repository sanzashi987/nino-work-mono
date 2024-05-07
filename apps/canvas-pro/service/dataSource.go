package service

import "context"

type DataSourceService struct{}

var DataSourceServiceImpl *DataSourceService = &DataSourceService{}

// data source

func (serv DataSourceService) ListDataSources(ctx context.Context) {

}
