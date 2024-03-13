package service

import (
	"context"
	"errors"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/apps/canvas-pro/enums"
	"github.com/cza14h/nino-work/pkg/db"
)

type ProjectService struct{}

var projectService *ProjectService

func init() {
	projectService = &ProjectService{}
}

func GetProjectService() *ProjectService {
	return projectService
}

func (p *ProjectService) Create(ctx context.Context, name, groupCode, jsonConfig, useTemplate string) (string, error) {
	projectDao := dao.NewProjectDao(ctx)

	newProject := &model.ProjectModel{}
	newProject.TypeTag, newProject.Name, newProject.Version, newProject.Config = enums.PROJECT, name, enums.DefaultVersion, jsonConfig

	if err := projectDao.Create(newProject); err != nil {
		return "", err
	}

	return newProject.Code, nil
}

func (p *ProjectService) Update(ctx context.Context, code string, name, config, thumbnail, group *string) (err error) {
	projectDao := dao.NewProjectDao(ctx)

	toUpdate, idProject := model.ProjectModel{}, model.ProjectModel{}
	var id uint64
	if id, _, err = enums.GetIdFromCode(code); err != nil {
		return
	}

	idProject.Id, toUpdate.Id = id, id
	projectDao.DB.First(&idProject)

	if name != nil {
		toUpdate.Name = *name
	}
	if thumbnail != nil {
		toUpdate.Thumbnail = *thumbnail
	}

	if config != nil {
		toUpdate.Config = *config
	}

	if group != nil {
		var groupId uint64
		if groupId, _, err = enums.GetIdFromCode(*group); err != nil {
			return
		}
		toUpdate.GroupId = groupId
	}

	return projectDao.UpdateById(toUpdate)
}

type ProjectDetail struct {
	Code       string
	Name       string
	Thumbnail  string
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

func (p *ProjectService) GetInfoById(ctx context.Context, code string) (result *ProjectDetail, err error) {
	projectDao := dao.NewProjectDao(ctx)
	project, e := projectDao.FindByKey("code", code)
	if e != nil {
		err = e
		return
	}

	result.Code, result.Name, result.Thumbnail = code, project.Name, project.Thumbnail
	result.CreateTime, result.UpdateTime = project.GetCreatedDate(), project.GetUpdatedDate()
	return
}

func (p *ProjectService) LogicalDeletion(ctx context.Context, codes []string) (err error) {
	projectDao := dao.NewProjectDao(ctx)

	intIds := []uint64{}

	for _, code := range codes {
		id, _, _ := enums.GetIdFromCode(code)
		intIds = append(intIds, id)
	}

	return projectDao.BatchLogicalDelete(intIds)
}

type ProjectInfo struct {
	Name      string
	Thumbnail string
	// Group     string
	db.BaseTime
}
type ProjectInfoList = []ProjectInfo

var ErrorUserWorkspaceNotMatch = errors.New("current user does not have the access right to the given workspace")

func (p *ProjectService) GetList(ctx context.Context, userId uint64, page, size int, workspace string, name, group *string) (*ProjectInfoList, error) {
	projectDao := dao.NewProjectDao(ctx)
	if !ValidateUserWorkspace(ctx, userId, workspace) {
		return nil, ErrorUserWorkspaceNotMatch
	}

	infos, err := projectDao.GetList(page, size, workspace, name, group)
	if err != nil {
		return nil, err
	}

	result := []ProjectInfo{}

	for _, info := range *infos {
		temp := ProjectInfo{}
		temp.Name, temp.CreateTime, temp.UpdateTime = info.Name, info.CreateTime, info.UpdateTime
		result = append(result, temp)
	}

	return &result, nil

}

func (p *ProjectService) PublishProject(ctx context.Context, Code string, PulishFlag int, PublishSecretKey *string) error {

}
