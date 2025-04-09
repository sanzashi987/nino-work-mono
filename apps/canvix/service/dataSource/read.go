package dataSource

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

type QuerySearchReq struct {
	SourceName string   `json:"sourceName"`
	SourceType []string `json:"sourceType"`
	Search     string   `json:"search"`
}

type ListReq struct {
	shared.PaginationRequest
	QuerySearchReq
}

func List(ctx context.Context, workspaceId uint64, payload *ListReq) ([]*DataSourceDetail, error) {
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

func GetDataSourceById(ctx context.Context, workspaceId uint64, sourceIdCode string) (*DataSourceDetail, error) {
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
