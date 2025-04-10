package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type GetWorkspaceInfoReq struct {
	WorkspaceCode string `json:"workspaceCode"`
}

type GroupInfo struct {
	Name string `json:"name"`
	Id   uint64 `json:"id"`
}

type GetWorkspaceInfoRes struct {
	Name   string       `json:"name"`
	Type   string       `json:"type"`
	Groups []*GroupInfo `json:"groups"`
}

func GetWorkspaceInfo(ctx *gin.Context, req *GetWorkspaceInfoReq) ([]*GetWorkspaceInfoRes, error) {
	var workspaceCode string = req.WorkspaceCode
	if workspaceCode == "" {
		
	}

	workspaceId, _, err := consts.GetIdFromCode(workspaceCode)
	if err != nil {
		return nil, err
	}

	tx := db.NewTx(ctx)

	allGroups := []*model.GroupModel{}

	if err := tx.Model(&model.GroupModel{}).Where("workspace = ?", workspaceId).Find(&allGroups).Error; err != nil {
		return nil, err
	}

	resMap := map[string]*GetWorkspaceInfoRes{}

	for _, group := range allGroups {
		if _, ok := resMap[group.TypeTag]; !ok {
			resMap[group.TypeTag] = &GetWorkspaceInfoRes{
				Name:   group.Name,
				Type:   consts.TagToName[group.TypeTag],
				Groups: []*GroupInfo{},
			}
		}

		resMap[group.TypeTag].Groups = append(resMap[group.TypeTag].Groups, &GroupInfo{
			Name: group.Name,
			Id:   group.Id,
		})
	}

	res := []*GetWorkspaceInfoRes{}
	for _, group := range resMap {
		res = append(res, group)
	}

	return res, nil
}
