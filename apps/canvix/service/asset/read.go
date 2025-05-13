package asset

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
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

type AssetInfo struct {
	FileCode   string  `json:"asset_code"`
	Name       string  `json:"asset_name"`
	MimeType   string  `json:"mime_type"`
	Size       int     `json:"size"`
	Suffix     *string `json:"suffix"`
	CreateTime string  `json:"create_time"`
	UpdateTime string  `json:"update_time"`
	TypeTag    string  `json:"type_tag"`
}

type ListAssetReq struct {
	GroupCode string `json:"group_code"`
	shared.PaginationRequest
}

type ListAssetResponse = shared.ResponseWithPagination[[]*AssetInfo]

func ListAssetByType(ctx context.Context, workspaceId uint64, typeTag string, req *ListAssetReq) (*ListAssetResponse, error) {
	var groupId *uint64
	if req.GroupCode != "" {
		if id, _, err := consts.GetIdFromCode(req.GroupCode); err != nil {
			return nil, err
		} else {
			groupId = &id
		}
	}

	tx := db.NewTx(ctx)

	condition := tx.Model(&model.AssetModel{}).Where("workspace = ? AND type_tag = ? AND group_id = ?", workspaceId, typeTag, *groupId)

	r, err := db.QueryWithTotal[model.AssetModel](condition, req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	data := []*AssetInfo{}
	for _, record := range r.Records {
		data = append(data, &AssetInfo{
			TypeTag:    typeTag,
			FileCode:   record.Code,
			Name:       record.Name,
			CreateTime: record.GetCreatedDate(),
			UpdateTime: record.GetUpdatedDate(),
		})
	}

	res := ListAssetResponse{}
	res.Init(data, r.Page, r.Total)

	return &res, nil
}
