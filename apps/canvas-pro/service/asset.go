package service

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
)

type AssetService struct{}

var AssetServiceImpl *AssetService = &AssetService{}

func (serv *AssetService) BatchMoveGroup(ctx context.Context, assetCodes []string, groupCode, workspaceCode string) error {
	workspaceId, groupId, projectIds, err := commonMoveGroup(assetCodes, groupCode, workspaceCode)
	if err != nil {
		return err
	}

	if err := dao.NewAssetDao(ctx).BatchMoveGroup(groupId, workspaceId, projectIds); err != nil {
		return err
	}

	return nil
}
