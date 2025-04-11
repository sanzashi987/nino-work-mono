package asset

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/shared"
	"github.com/sanzashi987/nino-work/proto/storage"
)

type AssetDetailResponse struct {
	Name      string `json:"fileName"`
	Code      string `json:"fileId"`
	GroupCode string `json:"groupCode"`
	MimeType  string `json:"mimeType"`
	Size      int64  `json:"size"`
	Suffix    string `json:"suffix"`
	shared.DBTime
}

func GetAssetDetail(ctx context.Context, uploadRpc storage.StorageService, workspaceId uint64, assetCode string) (*AssetDetailResponse, error) {
	tx := db.NewTx(ctx)
	assetId, _, _ := consts.GetIdFromCode(assetCode)

	record := model.AssetModel{}

	if err := tx.Where("id = ? AND workspace = ?", assetId, workspaceId).Find(&record).Error; err != nil {
		return nil, err
	}

	rpcReq := storage.FileQueryRequest{}
	rpcReq.Id = record.FileId

	rpcRes, err := uploadRpc.GetFileDetail(ctx, &rpcReq)
	if err != nil {
		return nil, err
	}

	result := AssetDetailResponse{
		Name:     record.Name,
		Code:     record.Code,
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

type ListAssetRes struct {
	FileCode   string  `json:"fileId"`
	Name       string  `json:"fileName"`
	MimeType   string  `json:"mimeType"`
	Size       int     `json:"size"`
	Suffix     *string `json:"suffix"`
	CreateTime string  `json:"createTime"`
	UpdateTime string  `json:"updateTime"`
}

type ListAssetReq struct {
	GroupCode string `json:"groupCode"`
	shared.PaginationRequest
}

func ListAssetByType(ctx context.Context, workspaceId uint64, typeTag string, payload *ListAssetReq) (int64, []*ListAssetRes, error) {
	var groupId *uint64
	if payload.GroupCode != "" {
		if id, _, err := consts.GetIdFromCode(payload.GroupCode); err != nil {
			return 0, nil, err
		} else {
			groupId = &id
		}
	}

	tx := db.NewTx(ctx)

	records, err := dao.ListAssets(tx, workspaceId, groupId, payload.Page, payload.Size, typeTag)
	if err != nil {
		return 0, nil, err
	}

	recordTotal, err := dao.GetAssetCount(tx, workspaceId, groupId, typeTag)
	if err != nil {
		return 0, nil, err
	}

	res := []*ListAssetRes{}
	for _, record := range records {
		res = append(res, &ListAssetRes{
			FileCode:   record.Code,
			Name:       record.Name,
			CreateTime: record.GetCreatedDate(),
			UpdateTime: record.GetUpdatedDate(),
		})
	}

	return recordTotal, res, nil
}
