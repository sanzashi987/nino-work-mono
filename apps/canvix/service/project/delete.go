package project

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

func Delete(ctx context.Context, workspaceId uint64, codes []string) error {
	tx := db.NewTx(ctx).Begin()

	intIds, err := consts.CodesIntoIds(codes)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&model.ProjectModel{}).Where("id IN ?", intIds).Delete(&model.ProjectModel{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
