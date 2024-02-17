package service

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
)

type ProjectService struct{}

var projectService *ProjectService

func init() {
	projectService = &ProjectService{}

}

func GetProjectService() *ProjectService {
	return projectService
}

func (p *ProjectService) Create(ctx context.Context, name, groupCode, jsonConfig, useTemplate string) {
	projectDao := dao.NewProjectDao(ctx)

	newProject := &model.ProjectModel{
		Name:    name,
		Version: "0.1.0",
		Config:  jsonConfig,
	}
	projectDao.Create(newProject)
}
