package project

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type ProjectUpdateReq struct {
	Code      string  `json:"code"`
	Name      *string `json:"name"`
	Config    *string `json:"rootConfig"`
	Thumbnail *string `json:"thumbnail"`
	GroupCode *string `json:"groupCode"`
}

func Update(ctx context.Context, workspaceId uint64, req *ProjectUpdateReq) error {
	tx := db.NewTx(ctx).Begin()

	code, name, config, thumbnail, group := req.Code, req.Name, req.Config, req.Thumbnail, req.GroupCode

	toUpdate, idProject := model.ProjectModel{}, model.ProjectModel{}

	id, _, err := consts.GetIdFromCode(code)
	if err != nil {
		return err
	}

	idProject.Id, toUpdate.Id = id, id
	if err := tx.First(&idProject).Error; err != nil {
		return err
	}

	if name != nil {

		err := consts.IsLegalName(*name)
		if err != nil {
			return err
		}

		toUpdate.Name = *name
	}
	if thumbnail != nil {
		toUpdate.Thumbnail = *thumbnail
	}

	if config != nil {
		toUpdate.Config = *config
	}

	if group != nil {
		groupId, _, err := consts.GetIdFromCode(*group)
		if err != nil {
			return err
		}

		// detect if the group is exist
		groupModel := model.GroupModel{}
		if err := tx.Where("id = ? AND workspace = ?", groupId, workspaceId).Find(&groupModel).Error; err != nil {
			return err
		}
		toUpdate.GroupId = groupId
	}

	if err := tx.Model(&idProject).Updates(&toUpdate).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
