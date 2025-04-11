package asset

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

func Delete(ctx context.Context, workspaceId uint64, codes []string) error {

	ids, err := consts.CodesIntoIds(codes)
	if err != nil {
		return err
	}

	tx := db.NewTx(ctx).Begin()

	if err := tx.Model(&model.AssetModel{}).Where("workspace = ? AND id in ?", workspaceId, ids).Delete(&model.ProjectModel{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
