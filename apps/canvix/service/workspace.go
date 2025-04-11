package service

import (
	"errors"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type GetConsolenfoReq struct {
	WorkspaceCode string `uri:"workspace_code"`
}

type GroupInfo struct {
	consts.CanvixCodeEnum
	Type string `json:"type"`
}

type WorkspaceInfo = consts.CanvixCodeEnum

type GetConsoleInfoRes struct {
	consts.CanvixCodeEnum
	Groups     []*GroupInfo     `json:"groups"`
	Workspaces []*WorkspaceInfo `json:"workspaces"`
}

var userHasNoWorkspace = errors.New("user does not belong to any workspace")

func GetConsoleInfo(ctx *gin.Context, req *GetConsolenfoReq) (*GetConsoleInfoRes, error) {
	workspaceCode := req.WorkspaceCode
	tx := db.NewTx(ctx)
	userId := ctx.GetUint64(controller.UserID)
	userModel, err := dao.GetUserWorkspaces(tx, userId)
	if err != nil {
		return nil, err
	}

	if workspaceCode == "" {

		if len(userModel.Workspaces) == 0 {
			return nil, userHasNoWorkspace
		}

		sort.SliceStable(userModel.Workspaces, func(i, j int) bool {
			return userModel.Workspaces[i].Id < userModel.Workspaces[j].Id
		})

		workspaceCode = userModel.Workspaces[0].Code
	}

	var activeWorkspace *model.WorkspaceModel
	for _, w := range userModel.Workspaces {
		if w.Code == workspaceCode {
			activeWorkspace = &w
			break
		}
	}

	if activeWorkspace == nil {
		return nil, userHasNoWorkspace
	}

	workspaceId, _, err := consts.GetIdFromCode(workspaceCode)
	if err != nil {
		return nil, err
	}

	allGroups := []*model.GroupModel{}

	if err := tx.Model(&model.GroupModel{}).Where("workspace = ?", workspaceId).Find(&allGroups).Error; err != nil {
		return nil, err
	}

	workspaces := []*WorkspaceInfo{}
	for _, workspace := range userModel.Workspaces {
		workspaces = append(workspaces, &WorkspaceInfo{
			Name: workspace.Name,
			Code: workspace.Code,
		})
	}

	groups := []*GroupInfo{}
	for _, group := range allGroups {
		g := &GroupInfo{Type: group.TypeTag}
		g.Name, g.Code = group.Name, group.Code
		groups = append(groups, g)
	}

	res := GetConsoleInfoRes{
		Groups:     groups,
		Workspaces: workspaces,
	}
	res.Name, res.Code = activeWorkspace.Name, activeWorkspace.Code

	return &res, nil
}
