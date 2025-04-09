package dataSource

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type UpdateReq struct {
	SourceName *string `json:"sourceName"`
	SourceInfo *string `json:"sourceInfo"`
	SourceId   string  `json:"sourceId" binding:"required"`
}

func Update(ctx context.Context, workspaceId uint64, payload *UpdateReq) (*DataSourceDetail, error) {
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
