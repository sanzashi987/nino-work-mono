package project

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type CreateProjectReq struct {
	Name string `json:"name" binding:"required"`
	// Version     string
	Config      string  `json:"root_config" binding:"required"`
	GroupCode   *string `json:"group_code"`
	UseTemplate *string `json:"template_id"` //template Id
}

func Create(ctx context.Context, workspaceId uint64, req *CreateProjectReq) (string, error) {
	tx := db.NewTx(ctx).Begin()
	name, jsonConfig, groupCode, useTemplate := req.Name, req.Config, req.GroupCode, req.UseTemplate
	newProject := &model.ProjectModel{}
	newProject.TypeTag, newProject.Name, newProject.Version, newProject.Config = consts.PROJECT, name, consts.DefaultVersion, jsonConfig
	newProject.Workspace = workspaceId
	if useTemplate != nil {
		templateId, _, _ := consts.GetIdFromCode(*useTemplate)

		template := model.TemplateModel{}

		if e := tx.Where("id = ?", templateId).Find(&template).Error; e != nil {
			return "", e
		}
		newProject.Config = template.Config
	}

	if groupCode != nil {
		groupId, _, err := consts.GetIdFromCode(*groupCode)
		if err != nil {
			return "", err
		}
		newProject.GroupId = groupId
	}

	if err := tx.Create(newProject).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	tx.Commit()
	return newProject.Code, nil
}

func Duplicate(ctx context.Context, workspaceId uint64, code string) (string, error) {
	tx := db.NewTx(ctx)

	copyFromId, _, err := consts.GetIdFromCode(code)
	if err != nil {
		return "", err
	}
	copyFrom := model.ProjectModel{}
	if err := tx.Where("id = ?", copyFromId).Find(&copyFrom).Error; err != nil {
		return "", err
	}
	name := copyFrom.Name + "_copied"
	hasGroupCode := copyFrom.GroupId == 0
	var groupCode *string = nil
	if hasGroupCode {
		str := consts.GetCodeFromId(consts.GROUP, copyFrom.GroupId)
		groupCode = &str
	}

	toCreate := CreateProjectReq{
		Name:      name,
		Config:    copyFrom.Config,
		GroupCode: groupCode,
	}

	return Create(ctx, workspaceId, &toCreate)
}
