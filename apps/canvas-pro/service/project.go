package service

import (
	"context"

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

type ProjectInfo struct {
	Name      string
	Thumbnail string
	// Group     string
	db.BaseTime
}
type ProjectListResponse = []ProjectInfo

func (p *ProjectService) GetList(ctx context.Context, page, size int, name, group, workspace string) (*ProjectListResponse, error) {
	projectDao := dao.NewProjectDao(ctx)

	infos, err := projectDao.GetList(page, size, name, group, workspace)
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
