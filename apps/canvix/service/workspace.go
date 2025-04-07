package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type GetWorkspaceInfoReq struct {
	WorkspaceCode string `json:"workspaceCode" binding:"required"`
}

type GroupInfo struct {
	Name   string `json:"name"`
	Id     uint64 `json:"id"`
	Custom bool   `json:"configurable"`
}

type GetWorkspaceInfoRes struct {
	Name   string       `json:"name"`
	Type   string       `json:"type"`
	Groups []*GroupInfo `json:"groups"`
}

func GetWorkspaceInfo(ctx *gin.Context, req *GetWorkspaceInfoReq) (*GetWorkspaceInfoRes, error) {
	workspaceId, _, err := consts.GetIdFromCode(req.WorkspaceCode)
	if err != nil {
		return nil, err
	}

	tx := db.NewTx(ctx)

	allGroups := []*model.GroupModel{}

	if err := tx.Model(&model.GroupModel{}).Where("workspace = ?", workspaceId).Find(&allGroups).Error; err != nil {

		return nil, err
	}

	return nil, nil
}
