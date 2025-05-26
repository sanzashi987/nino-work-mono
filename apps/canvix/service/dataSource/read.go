package dataSource

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/shared"
	"gorm.io/gorm"
)

type QuerySearchReq struct {
	SourceName string   `json:"source_name"`
	SourceType []string `json:"source_type"`
	Search     string   `json:"search"`
}

type ListReq struct {
	shared.PaginationRequest
	QuerySearchReq
}

type ListRes = shared.ResponseWithPagination[[]*DataSourceDetail]

func getListCommonQuery(tx *gorm.DB, workspaceId uint64, payload *ListReq) *gorm.DB {
	query := tx.Model(&model.DataSourceModel{}).Where("workspace = ?", workspaceId)

	if payload.SourceName != "" {
		query = query.Where("name LIKE ?", payload.SourceName)
	}

	if len(payload.SourceType) != 0 {
		sourceTypeEnums := []uint8{}

		for _, v := range payload.SourceType {
			enum, exist := model.SourceTypeStringToEnum[v]
			if exist {
				sourceTypeEnums = append(sourceTypeEnums, enum)
			}
		}

		query = query.Where("source_type IN ?", sourceTypeEnums)
	}
	return query
}

func List(ctx context.Context, workspaceId uint64, payload *ListReq) (*ListRes, error) {
	tx := db.NewTx(ctx)

	query := getListCommonQuery(tx, workspaceId, payload)

	r, err := db.QueryWithTotal[model.DataSourceModel](query, payload.Page, payload.Size)
	if err != nil {
		return nil, err
	}

	data := []*DataSourceDetail{}
	for _, source := range r.Records {
		temp := intoDataSourceDetail(source)

		data = append(data, temp)
	}

	res := ListRes{}
	res.Init(data, r.Page, r.Total)

	return &res, nil
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
