package service

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
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

type ListAssetReq struct {
	GroupCode string `json:"groupCode"`
	// Name      string `json:"fileName"`
	shared.PaginationRequest
}

func (serv *AssetService) ListAssetByType(ctx context.Context, workspaceId uint64, typeTag string, payload *ListAssetReq) (int64, []*ListAssetRes, error) {

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

type GroupCount struct {
	Id    uint64 `gorm:"column:id"`
	Count uint64 `gorm:"column:count"`
}

func (serv AssetService) GetCountFromGroupId(ctx context.Context, workspaceId uint64, groupId []uint64) ([]*GroupCount, error) {
	tx := db.NewTx(ctx)
	res := []*GroupCount{}
	var assetTableName = model.ProjectModel{}.TableName()
	err := tx.Table(assetTableName).Where("workspace = ?", workspaceId).Where("group_id IN ?", groupId).Select("id", "COUNT(id) as count").Group("group_id").Find(&res).Error
	return res, err
}

func (serv *AssetService) BatchMoveGroup(ctx context.Context, workspaceId uint64, assetCodes []string, groupName, groupCode string) error {
	code := groupCode
	tx := db.NewTx(ctx).Begin()

	if newGroup, err := createGroup(tx, workspaceId, groupName, consts.DESIGN); err != nil {
		return err
	} else if newGroup != nil {
		code = newGroup.Code
	}

	groupId, projectIds, err := commonMoveGroup(assetCodes, code)
	if err != nil {
		return err
	}

	if dao.AssetBatchMoveGroup(tx, groupId, workspaceId, projectIds); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

const chunkSize = 1024 * 1024 / 2

func (serv AssetService) UpdateName(ctx context.Context, workspaceId uint64, assetName, assetCode string) error {
	if err := consts.IsLegalName(assetName); err != nil {
		return err
	}
	assetId, _, _ := consts.GetIdFromCode(assetCode)

	tx := db.NewTx(ctx)

	return dao.UpdateAssetName(tx, workspaceId, assetId, assetName)

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

func DeleteAssets(ctx context.Context, workspaceId uint64, assetCode []string) error {
	// if err := consts.IsLegalName(assetName); err != nil {
	// 	return err
	// }
	// assetId, _, _ := consts.GetIdFromCode(assetCode)

	// tx := db.NewTx(ctx)

	// return dao.UpdateAssetName(tx, workspaceId, assetId, assetName)

}
