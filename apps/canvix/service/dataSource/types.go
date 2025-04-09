package dataSource

import (
	"strconv"

	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

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
