package service

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
	"github.com/cza14h/nino-work/proto/upload"
)

type AssetService struct{}

var AssetServiceImpl *AssetService = &AssetService{}

func (serv *AssetService) BatchMoveGroup(ctx context.Context, workspaceId uint64, assetCodes []string, groupCode string) error {
	groupId, projectIds, err := commonMoveGroup(assetCodes, groupCode)
	if err != nil {
		return err
	}

	if err := dao.NewAssetDao(ctx).BatchMoveGroup(groupId, workspaceId, projectIds); err != nil {
		return err
	}

	return nil
}

const chunkSize = 1024 * 1024

func (serv *AssetService) UploadFile(ctx context.Context, uploadRpc upload.FileUploadService, workspaceId uint64, groupName, groupCode, filename, assetType string, file *multipart.FileHeader) (err error) {
	stream, err := uploadRpc.UploadFile(ctx)
	if err != nil {
		return err
	}

	reader, _ := file.Open()
	defer reader.Close()
	defer stream.Close()
	for {
		var n int
		buf := make([]byte, chunkSize)
		n, err = reader.Read(buf)
		if err = stream.Send(&upload.FileUploadRequest{
			Filename: filename,
			Data:     buf[:n],
		}); err != nil {
			return
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}
	}

	if err = stream.CloseSend(); err != nil {
		return
	}
	rpcResponse := upload.FileUploadResponse{}
	if err = stream.RecvMsg(&rpcResponse); err != nil {
		return
	}

	assetDao := dao.NewAssetDao(ctx)

	if err = assetDao.CreateAsset(workspaceId, rpcResponse.Id, assetType); err != nil {
		return
	}

	return
}
