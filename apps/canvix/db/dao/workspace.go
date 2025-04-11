package dao

import (
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
)

func CreateWorkspace(canvasUser *model.CanvixUserModel) {
	userId := canvasUser.UnifiedUserId
	newWorkspace := model.WorkspaceModel{Owner: userId}
	newWorkspace.Creator = userId

}
