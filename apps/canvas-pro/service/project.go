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

func (p *ProjectService) Update(ctx context.Context, code, name string) {
	projectDao := dao.NewProjectDao(ctx)

}

type ProjectDetailResponse struct {
	Id   string
	Code string
	db.BaseTime
}

func (p *ProjectService) GetInfoById(ctx context.Context, id string) (*ProjectDetailResponse, error) {
	projectDao := dao.NewProjectDao(ctx)
	result := &ProjectDetailResponse{}
	project, err := projectDao.FindByKey("id", id)
	if err != nil {
		return nil, err
	}

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
