package dataSource

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type CreateReq struct {
	SourceName string `json:"source_name" binding:"required"`
	SourceType string `json:"source_type" binding:"required"`
	SourceInfo string `json:"source_info" binding:"required"`
}

func Create(ctx context.Context, workspaceId uint64, payload *CreateReq) (*DataSourceDetail, error) {
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
