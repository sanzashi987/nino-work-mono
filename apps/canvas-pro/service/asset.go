package service

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/cza14h/nino-work/apps/canvas-pro/consts"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/proto/upload"
)

type AssetService struct{}

var AssetServiceImpl *AssetService = &AssetService{}

func (serv *AssetService) BatchMoveGroup(ctx context.Context, workspaceId uint64, assetCodes []string, groupName, groupCode string) error {
	code := groupCode
	assetDao := dao.NewAssetDao(ctx)
	assetDao.BeginTransaction()

	if newGroup, err := createGroup(ctx, (*dao.AnyDao[model.AssetModel])(assetDao), workspaceId, groupName, consts.DESIGN); err != nil {
		return err
	} else if newGroup != nil {
		code = newGroup.Code
	}

	groupId, projectIds, err := commonMoveGroup(assetCodes, code)
	if err != nil {
		return err
	}

	if assetDao.BatchMoveGroup(groupId, workspaceId, projectIds); err != nil {
		assetDao.RollbackTransaction()
		return err
	}

	assetDao.CommitTransaction()

	return nil
}

const chunkSize = 1024 * 1024

type UploadAssetRes struct {
	FileId   string `json:"fileId"`
	MimeType string `json:"mimeType"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Suffix   string `json:"suffix"`
}

func (serv *AssetService) UploadFile(ctx context.Context, uploadRpc upload.FileUploadService, workspaceId uint64, groupName, groupCode, filename, typeTag string, file *multipart.FileHeader) (res *UploadAssetRes, err error) {
	stream, err := uploadRpc.UploadFile(ctx)
	if err != nil {
		return
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
	assetDao.BeginTransaction()
	code := groupCode
	if newGroup, err := createGroup(ctx, (*dao.AnyDao[model.AssetModel])(assetDao), workspaceId, groupName, consts.DESIGN); err != nil {
		return nil, err
	} else if newGroup != nil {
		code = newGroup.Code
	}

	groupId, typeTag, err := consts.GetIdFromCode(code)
	if !consts.IsGroup(typeTag) {
		return //nil, errors.New("not a group tag")
	}

	if err != nil {
		return
	}

	asset, err := assetDao.CreateAsset(workspaceId, groupId, filename, rpcResponse.Id, typeTag)

	if err != nil {
		assetDao.RollbackTransaction()
		return
	}

	assetDao.CommitTransaction()
	// TODO size not get
	res.Suffix, res.Name, res.MimeType, res.FileId = rpcResponse.Extension, asset.Name, rpcResponse.MimeType, asset.Code
	return
}
