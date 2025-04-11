package asset

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type UpdateAssetReq struct {
	FileId   string `json:"fileId" binding:"required"`
	FileName string `json:"fileName" binding:"required"`
}

const chunkSize = 1024 * 1024 / 2

func UpdateName(ctx context.Context, workspaceId uint64, req *UpdateAssetReq) error {
	if err := consts.IsLegalName(req.FileName); err != nil {
		return err
	}
	assetId, _, _ := consts.GetIdFromCode(req.FileId)

	tx := db.NewTx(ctx).Begin()

	if err := tx.Model(&model.AssetModel{}).Where("id = ? AND workspace = ?", assetId, workspaceId).Update("name", req.FileName).Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
