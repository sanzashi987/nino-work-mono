package service

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/consts"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/model"
	"github.com/sanzashi987/nino-work/pkg/shared"
	"github.com/sanzashi987/nino-work/proto/storage"
)

type AssetService struct{}

var AssetServiceImpl *AssetService = &AssetService{}

type ListAssetRes struct {
	FileCode string `json:"fileId"`
	Name     string `json:"fileName"`
	// GroupCode  string  `json:"groupCode"`
	// GroupName  string  `json:"groupName"`
	MimeType   string  `json:"mimeType"`
	Size       int     `json:"size"`
	Suffix     *string `json:"suffix"`
	CreateTime string  `json:"createTime"`
	UpdateTime string  `json:"updateTime"`
}

func (serv *AssetService) ListAssetByType(ctx context.Context, workspaceId uint64, page, size int, typeTag, groupCode string) (recordTotal int64, res []ListAssetRes, err error) {

	var groupId *uint64
	if groupCode != "" {

		if id, _, e := consts.GetIdFromCode(groupCode); e != nil {
			err = e
			return
		} else {
			groupId = &id
		}
	}

	assetDao := dao.NewAssetDao(ctx)
	records, err := assetDao.ListAssets(workspaceId, groupId, page, size, typeTag)
	if err != nil {
		return
	}

	recordTotal, err = assetDao.GetAssetCount(workspaceId, groupId, page, size, typeTag)
	if err != nil {
		return
	}

	for _, record := range records {
		res = append(res, ListAssetRes{
			FileCode:   record.Code,
			Name:       record.Name,
			CreateTime: record.GetCreatedDate(),
			UpdateTime: record.GetUpdatedDate(),
		})
	}

	return

}

func (serv AssetService) GetCountFromGroupId(ctx context.Context, workspaceId uint64, groupId []uint64) ([]dao.GroupCount, error) {
	assetDao := dao.NewAssetDao(ctx)

	return assetDao.GetAssetCountByGroup(workspaceId, groupId)
}

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

const chunkSize = 1024 * 1024 / 2

type UploadAssetResponse struct {
	FileId   string `json:"fileId"`
	MimeType string `json:"mimeType"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Suffix   string `json:"suffix"`
}

func (serv *AssetService) UploadFile(ctx context.Context, uploadRpc storage.StorageService, workspaceId uint64, groupName, groupCode, filename, typeTag string, file *multipart.FileHeader) (res *UploadAssetResponse, err error) {
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
		if err = stream.Send(&storage.FileUploadRequest{
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
	rpcResponse := storage.FileDetailResponse{}
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
	res.Size, res.FileId = rpcResponse.Size, asset.Code
	res.Suffix, res.Name, res.MimeType = rpcResponse.Extension, asset.Name, rpcResponse.MimeType
	return
}

func (serv AssetService) UpdateName(ctx context.Context, workspaceId uint64, assetName, assetCode string) error {
	if err := consts.IsLegalName(assetName); err != nil {
		return err
	}

	assetId, _, _ := consts.GetIdFromCode(assetCode)

	assetDao := dao.NewAssetDao(ctx)

	return assetDao.UpdateAssetName(workspaceId, assetId, assetName)

}

type AssetDetailResponse struct {
	Name      string `json:"fileName"`
	Code      string `json:"fileId"`
	GroupCode string `json:"groupCode"`

	MimeType string `json:"mimeType"`
	Size     int64  `json:"size"`
	Suffix   string `json:"suffix"`

	shared.DBTime
}

func (serv AssetService) GetAssetDetail(ctx context.Context, uploadRpc storage.StorageService, workspaceId uint64, assetCode string) (*AssetDetailResponse, error) {

	assetDao := dao.NewAssetDao(ctx)
	assetId, _, _ := consts.GetIdFromCode(assetCode)

	record, err := assetDao.GetSingleAsset(workspaceId, assetId)
	if err != nil {
		return nil, err
	}

	rpcReq := storage.FileQueryRequest{}
	rpcReq.Id = record.FileId

	rpcRes, err := uploadRpc.GetFileDetail(ctx, &rpcReq)
	if err != nil {
		return nil, err
	}
	// result.CreateTime = record.

	result := AssetDetailResponse{
		Name: record.Name,
		Code: record.Code,
		// GroupCode: record.GroupId,

		MimeType: rpcRes.MimeType,
		Size:     rpcRes.Size,
		Suffix:   rpcRes.Extension,
		DBTime: shared.DBTime{
			CreateTime: record.GetCreatedDate(),
			UpdateTime: record.GetUpdatedDate(),
		},
	}

	return &result, err
}
