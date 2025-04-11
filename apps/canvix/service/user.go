package service

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/pkg/db"
)

func ValidateUserWorkspace(ctx context.Context, userId uint64, workspaceCode string) bool {
	tx := db.NewTx(ctx)

	user, err := dao.GetUserWorkspaces(tx, userId)
	if err != nil {
		return false
	}

	for _, workspace := range user.Workspaces {
		if workspace.Code == workspaceCode {
			return true
		}
	}
	return false
}

func UserOnBoard(ctx context.Context, userId uint64) {
	tx := db.NewTx(ctx).Begin()

}
